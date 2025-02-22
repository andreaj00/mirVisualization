/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package crypto provides an implementation of the MirModule module.
// It supports RSA and ECDSA signatures.
package crypto

import (
	cstd "crypto"
	"crypto/ecdsa"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
	"strings"

	es "github.com/go-errors/errors"

	t "github.com/filecoin-project/mir/pkg/types"
)

// DefaultImpl represents a generic implementation of the MirModule module that can be used at Node instantiation
// (when calling mir.NewNode)
type DefaultImpl struct {

	// Private key used for signing.
	privKey interface{}

	// Node public keys used for verifying signatures.
	nodeKeys map[t.NodeID]interface{}
}

// NewDefaultImpl returns a new initialized instance of a MirModule implementation.
// privKey is the serialized representation of the private key that will be used for signing.
// privKey must be the output of SerializePrivKey or GenerateKeyPair.
func NewDefaultImpl(privKey []byte) (*DefaultImpl, error) {

	// Deserialize the passed private key.
	key, err := privKeyFromBytes(privKey)
	if err != nil {
		// Report error if deserialization of the private key fails.
		return nil, es.Errorf("error parsing private key: %w", err)
	}

	// If deserialization succeeds, return the pointer to a new initialized instance of MirModule.
	return &DefaultImpl{
		privKey:  key,
		nodeKeys: make(map[t.NodeID]interface{}),
	}, nil
}

// Sign signs the provided data and returns the resulting signature.
// First, Sign computes a SHA256 hash of the concatenation of all the byte slices in data.
// Then it signs the hash using the private key specified at creation of this MirModule object.
func (c *DefaultImpl) Sign(data [][]byte) ([]byte, error) {
	switch key := c.privKey.(type) {
	case *rsa.PrivateKey:
		return key.Sign(crand.Reader, digest(data), cstd.SHA256)
	case *ecdsa.PrivateKey:
		return signEcdsa(key, digest(data))
	default:
		return nil, es.Errorf("unsupported private key type: %T", key)
	}
}

// RegisterNodeKey associates a public key with a node ID.
// pubKey must be the output of SerializePubKey.
// Calls to Verify will fail until RegisterNodeKey is successfully called with the corresponding node ID.
// Returns nil on success, a non-nil error on failure.
func (c *DefaultImpl) RegisterNodeKey(pubKey []byte, nodeID t.NodeID) error {

	// Deserialize passed public key
	key, err := pubKeyFromBytes(pubKey)
	if err != nil {
		// If deserialization fails, report error.
		return es.Errorf("error parsing node public key: %w", err)
	}

	// If deserialization succeeds, save public key under the given node ID.
	c.nodeKeys[nodeID] = key

	return nil
}

// DeleteNodeKey removes the public key associated with nodeID from the internal state.
// Any subsequent call to Verify(..., nodeID) will fail.
func (c *DefaultImpl) DeleteNodeKey(nodeID t.NodeID) {
	delete(c.nodeKeys, nodeID)
}

// Verify verifies a signature produced by the node with ID nodeID over data.
// First, Verify computes a SHA256 hash of the concatenation of all the byte slices in data.
// Then it verifies the signature over this hash using the public key registered under nodeID.
// Returns nil on success (i.e., if the given signature is valid) and a non-nil error otherwise.
// Note that RegisterNodeKey must be used to register the node's public key before calling Verify,
// otherwise Verify will fail.
func (c *DefaultImpl) Verify(data [][]byte, signature []byte, nodeID t.NodeID) error {

	pubKey, ok := c.nodeKeys[nodeID]
	if !ok {
		return es.Errorf("no public key for node with ID %v", nodeID)
	}

	return c.verifySig(data, signature, pubKey)
}

// verifySig performs the actual signature verification.
// It is called by Verify after looking up the appropriate verification key.
func (c *DefaultImpl) verifySig(data [][]byte, signature []byte, pubKey interface{}) error {
	if signature == nil {
		return es.Errorf("signature is nil")
	}

	switch key := pubKey.(type) {
	case *ecdsa.PublicKey:
		return verifyEcdsaSignature(key, digest(data), signature)
	case *rsa.PublicKey:
		return rsa.VerifyPKCS1v15(key, cstd.SHA256, digest(data), signature)
	default:
		return es.Errorf("unsupported public key type: %T", key)
	}
}

// GenerateKeyPair generates a pair of ECDSA keys that can be used for signing and verifying.
// The randomness parameter should be backed by a high-quality source of entropy such as crypto/rand.Reader.
// The priv key can be used for creation of a new instance of the crypto module (New function)
// and the pub key can be passed to DefaultImpl.RegisterNodeKey.
func GenerateKeyPair(randomness io.Reader) (priv []byte, pub []byte, err error) {

	// Generate ECDSA keys.
	var privKey, pubKey interface{}
	privKey, pubKey, err = generateEcdsaKeyPair(randomness)
	if err != nil {
		return nil, nil, es.Errorf("error generating ecdsa key: %w", err)
	}

	// Serialize private key.
	if priv, err = SerializePrivKey(privKey); err != nil {
		return nil, nil, es.Errorf("error serializing private key: %w", err)
	}

	// Serialized public key.
	if pub, err = SerializePubKey(pubKey); err != nil {
		return nil, nil, es.Errorf("error serializing public key: %w", err)
	}

	// All output variables have been set, just return.
	return
}

// digest computes the SHA256 of the concatenation of all byte slices in data.
func digest(data [][]byte) []byte {
	h := sha256.New()
	for _, d := range data {
		h.Write(d)
	}
	return h.Sum(nil)
}

// PrivKeyFromBytes deserializes a private key returned by GenerateKeyPair or SerializePrivKey.
func privKeyFromBytes(raw []byte) (interface{}, error) {

	// Parse key from raw bytes.
	pk, err := x509.ParsePKCS8PrivateKey(raw)
	if err != nil {
		return nil, err
	}

	// Check if key type is supported.
	switch p := pk.(type) {
	case *ecdsa.PrivateKey, *rsa.PrivateKey:
		return p, nil
	default:
		return nil, es.Errorf("unsupported private key type: %T", p)
	}
}

// PubKeyFromBytes deserializes a public key returned by GenerateKeyPair or SerializePubKey.
func pubKeyFromBytes(raw []byte) (interface{}, error) {

	// Parse key from raw bytes.
	pk, err := x509.ParsePKIXPublicKey(raw)
	if err != nil {
		return nil, err
	}

	// Check if key type is supported.
	switch p := pk.(type) {
	case *ecdsa.PublicKey, *rsa.PublicKey:
		return p, nil
	default:
		return nil, es.Errorf("unsupported public key type: %T", p)
	}
}

// SerializePubKey serializes a public key into a byte slice.
// The output of this function can be used with DefaultImpl.RegisterNodeKey.
// Currently, pointers to crypto/ecdsa.PublicKey and crypto/rsa.PublicKey are supported types of pubKey.
func SerializePubKey(pubKey interface{}) (pubKeyBytes []byte, err error) {

	// Check if the key is of one of the supported types.
	switch key := pubKey.(type) {
	case *ecdsa.PublicKey, *rsa.PublicKey:
		// Serialize key if supported.
		return x509.MarshalPKIXPublicKey(key)
	default:
		// Return error if key type is not supported.
		return nil, es.Errorf("unsupported public key type: %T", key)
	}
}

// SerializePrivKey serializes a private key into a byte slice.
// The output of this function can be passed to New when creating an instance of DefaultImpl.
// Currently, pointers to crypto/ecdsa.PrivateKey and crypto/rsa.PrivateKey are supported types of pubKey.
func SerializePrivKey(privKey interface{}) (privKeyBytes []byte, err error) {

	// Check if the key is of one of the supported types.
	switch key := privKey.(type) {
	case *ecdsa.PrivateKey, *rsa.PrivateKey:
		// Serialize key if supported.
		return x509.MarshalPKCS8PrivateKey(key)
	default:
		// Return error if key type is not supported.
		return nil, es.Errorf("unsupported private key type: %T", key)
	}
}

// PubKeyFromFile extracts a public key from a PEM certificate file.
// Returns a serialized form of the public key
// that can be used directly with DefaultImpl.RegisterNodeKey.
func PubKeyFromFile(fileName string) ([]byte, error) {

	// Read contents of the file.
	certBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Decode the data.
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, es.Errorf("failed to decode PEM block")
	}

	// If the file only contains the key, return the bytes of the key directly.
	if block.Type == "PUBLIC KEY" {
		return block.Bytes, nil
	}

	// If the file is a certificate, parse out the key.
	if block.Type == "CERTIFICATE" {

		// Parse the certificate.
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}

		// Check if the key is of one of the supported types.
		switch key := cert.PublicKey.(type) {
		case *ecdsa.PublicKey, *rsa.PublicKey:
			// Return serialized key if supported.
			return SerializePubKey(key)
		default:
			// Return error if key type is not supported.
			return nil, es.Errorf("unsupported public key type: %T", key)
		}
	}

	// Return error if public key was not found in the file.
	return nil, es.Errorf("failed to find public key in the PEM block")
}

// PrivKeyFromFile extracts a private key from a PEM key file.
// Returns a serialized form of the private key.
// The output of this function can be passed to New when creating an instance of DefaultImpl.
func PrivKeyFromFile(file string) ([]byte, error) {

	// Read contents of the file.
	fileData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Loop over all PEM blocks in the file, starting with the first one.
	block, rest := pem.Decode(fileData)
	for block != nil {

		if key, err := privKeyFromPEMBlock(block); err == nil {
			// When a block with a private key is found, return it.
			return SerializePrivKey(key)
		}

		// Otherwise, try next block.
		block, rest = pem.Decode(rest)
	}

	return nil, es.Errorf("no valid key PEM block found")
}

// privKeyFromPEMBlock extracts a private key from block of a PEM file.
// If the block does not contain a private key, returns nil as the key and a corresponding non-nil error.
func privKeyFromPEMBlock(block *pem.Block) (interface{}, error) {
	if block == nil {
		return nil, es.Errorf("PEM block is nil")
	}

	if !strings.Contains(block.Type, "PRIVATE KEY") {
		return nil, es.Errorf("wrong PEM block type: %s", block.Type)
	}

	return privKeyFromBytes(block.Bytes)
}
