/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// ********************************************************************************
//                                                                               //
//         Chat demo application for demonstrating the usage of Mir              //
//                             (main executable)                                 //
//                                                                               //
// ********************************************************************************

package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	es "github.com/go-errors/errors"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/filecoin-project/mir"
	"github.com/filecoin-project/mir/pkg/checkpoint"
	mircrypto "github.com/filecoin-project/mir/pkg/crypto"
	"github.com/filecoin-project/mir/pkg/eventlog"
	"github.com/filecoin-project/mir/pkg/events"
	lsp "github.com/filecoin-project/mir/pkg/iss/leaderselectionpolicy"
	"github.com/filecoin-project/mir/pkg/logging"
	"github.com/filecoin-project/mir/pkg/membership"
	libp2p2 "github.com/filecoin-project/mir/pkg/net/libp2p"
	mempoolpbevents "github.com/filecoin-project/mir/pkg/pb/mempoolpb/events"
	trantorpbtypes "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	"github.com/filecoin-project/mir/pkg/trantor"
	tt "github.com/filecoin-project/mir/pkg/trantor/types"
	t "github.com/filecoin-project/mir/pkg/types"
	"github.com/filecoin-project/mir/pkg/util/errstack"
	"github.com/filecoin-project/mir/pkg/util/libp2p"
)

// parsedArgs represents parsed command-line parameters passed to the program.
type parsedArgs struct {

	// ID of this node.
	// The package github.com/hyperledger-labs/mir/pkg/types defines this and other types used by the library.
	OwnID t.NodeID

	// If set, print debug output to stdout.
	Verbose bool

	// If set, print trace output to stdout.
	Trace bool

	// Network transport type
	Net string

	// Name of the file containing the initial membership for joining nodes.
	InitMembershipFile string

	// Name of the file containing the initial state checkpoint to start from.
	InitChkpFile string

	// Name of the directory to store checkpoint files in.
	ChkpDir string
}

// main is just the wrapper for executing the run() and printing a potential error.
func main() {
	if err := run(); err != nil {
		errstack.Println(err)
		os.Exit(1)
	}
}

// run is the actual code of the program.
func run() error {

	// ================================================================================
	// Basic initialization and configuration
	// ================================================================================

	// Convenience variables
	var err error
	ctx := context.Background()

	// Parse command-line parameters.
	args := parseArgs(os.Args)

	// Initialize logger that will be used throughout the code to print log messages.
	var logger logging.Logger
	if args.Trace {
		logger = logging.ConsoleTraceLogger // Print trace-level info.
	} else if args.Verbose {
		logger = logging.ConsoleDebugLogger // Print debug-level info in verbose mode.
	} else {
		logger = logging.ConsoleWarnLogger // Only print errors and warnings by default.
	}

	fmt.Println("Initializing...")

	// ================================================================================
	// Load system membership info: IDs, addresses, ports, etc...
	// ================================================================================

	// For the dummy chat application, we require node IDs to be numeric,
	// as other metadata is derived from node IDs.
	ownNumericID, err := strconv.Atoi(string(args.OwnID))
	if err != nil {
		return errors.Wrap(err, "node IDs must be numeric in the sample app")
	}

	// Load initial system membership from the file indicated through the command line.
	initialAddrs, err := membership.FromFileName(args.InitMembershipFile)
	if err != nil {
		return errors.Wrap(err, "could not load membership")
	}
	initialMembership, err := membership.DummyMultiAddrs(initialAddrs)
	if err != nil {
		return errors.Wrap(err, "could not create dummy multiaddrs")
	}

	// ================================================================================
	// Instantiate the Mir node with the appropriate set of modules.
	// ================================================================================

	// Assemble listening address.
	// In this demo code, we always listen on tha address 0.0.0.0.
	portStr, err := getPortStr(initialMembership.Nodes[args.OwnID].Addr)
	if err != nil {
		return es.Errorf("could not parse port from own address: %w", err)
	}
	addrStr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", portStr)
	listenAddr, err := multiaddr.NewMultiaddr(addrStr)
	if err != nil {
		return es.Errorf("could not create listen address: %w", err)
	}

	// We use the default SMR parameters. The initial membership is, regardless of the starting checkpoint,
	// always the very first membership at sequence number 0. It is part of the system configuration.
	trantorParams := trantor.DefaultParams(initialMembership)
	trantorParams.Iss.LeaderSelectionPolicy = lsp.Simple

	// Create a dummy libp2p host for network communication (this is why we need a numeric ID)
	h, err := libp2p.NewDummyHostWithPrivKey(
		t.NodeAddress(libp2p.NewDummyMultiaddr(ownNumericID, listenAddr)),
		libp2p.NewDummyHostKey(ownNumericID),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create libp2p host")
	}

	// Initialize the libp2p transport subsystem.
	transport := libp2p2.NewTransport(trantorParams.Net, args.OwnID, h, logger, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create libp2p transport")
	}

	// Create a dummy crypto implementation that locally generates all keys in a pseudo-random manner.
	// localCrypto := deploytest.NewLocalCryptoSystem("pseudo", membership.GetIDs(initialMembership), logger)
	crypto := &mircrypto.DummyCrypto{DummySig: []byte{0}}

	// Assemble checkpoint directory name and instantiate the chat app logic.
	chkpDir := ""
	if args.ChkpDir != "" {
		chkpDir = args.ChkpDir + "/" + string(args.OwnID)
	}
	chatApp := NewChatApp(initialMembership, chkpDir)

	// genesis is s stable checkpoint (as given to the app's Checkpoint callback)
	// defining the initial state and configuration of the system.
	var genesis *checkpoint.StableCheckpoint

	if args.InitChkpFile != "" {

		// Load starting checkpoint from file if given.
		genesis, err = loadStableCheckpoint(args.InitChkpFile)
		if err != nil {
			return errors.Wrap(err, "could not load starting checkpoint from file")
		}

		// Verify that the starting checkpoint is valid.
		// err = genesis.VerifyCert(crypto.SHA256, localCrypto.Crypto(args.OwnID), trantorParams.Iss.InitialMembership)
		// if err != nil {
		//	return errors.Wrap(err, "starting checkpoint invalid")
		// }

	} else {
		// If no starting checkpoint is given, we create a new one from the initial membership.
		initialSnapshot, err := chatApp.Snapshot()
		if err != nil {
			return errors.Wrap(err, "could not create initial snapshot")
		}
		genesis, err = trantor.GenesisCheckpoint(initialSnapshot, trantorParams)
		if err != nil {
			return errors.Wrap(err, "could not create genesis checkpoint")
		}
	}

	// Create a Mir SMR system.
	trantorSystem, err := trantor.New(
		args.OwnID,
		transport,
		genesis,
		crypto,
		chatApp,
		trantorParams,
		logger,
	)
	if err != nil {
		return errors.Wrap(err, "could not create SMR system")
	}

	ownIDInt, _ := strconv.Atoi(string(args.OwnID))

	// Initialize recording of events
	interceptor, err := eventlog.NewRecorder(
		args.OwnID,
		fmt.Sprintf("node%d", ownIDInt),
		logging.Decorate(logging.ConsoleTraceLogger, "Interceptor: "),
		eventlog.FileSplitterOpt(eventlog.EventNewEpochLogger(trantor.DefaultModuleConfig().BatchFetcher)),
	)
	if err != nil {
		return errors.Wrap(err, "could not create new recorder")
	}
	// Create a Mir node, passing it all the modules of the SMR system.
	node, err := mir.NewNode(args.OwnID, mir.DefaultNodeConfig().WithLogger(logger), trantorSystem.Modules(), interceptor)
	if err != nil {
		return errors.Wrap(err, "could not create node")
	}

	// ================================================================================
	// Start the Node by launching necessary processing threads.
	// ================================================================================

	// Start the SMR system.
	// This will start all the goroutines that need to run within the modules of the SMR system.
	// For example, the network module will start listening for incoming connections and create outgoing ones.
	// The modules will become ready to be used by the node (but the node itself is not yet started).
	if err := trantorSystem.Start(); err != nil {
		return errors.Wrap(err, "could not start SMR system")
	}

	// Start the node in a separate goroutine
	nodeErr := make(chan error) // The error returned from running the Node will be written here.
	go func() {
		nodeErr <- node.Run(ctx)
	}()

	// ================================================================================
	// Read chat messages from stdin and submit them as transactions.
	// ================================================================================

	scanner := bufio.NewScanner(os.Stdin)

	// Prompt for chat message input.
	fmt.Println("Type in your messages and press 'Enter' to send.")

	// Read chat message from stdin.
	nextTxNo := tt.TxNo(0)
	for scanner.Scan() {

		// Submit the chat message as transaction payload to the mempool module.
		err := node.InjectEvents(ctx, events.ListOf(mempoolpbevents.NewTransactions(
			"mempool",
			[]*trantorpbtypes.Transaction{{
				ClientId: tt.ClientID(args.OwnID),
				TxNo:     nextTxNo,
				Type:     0,
				Data:     scanner.Bytes(),
			}}).Pb()),
		)

		// Print error if occurred.
		if err != nil {
			fmt.Println(err)
		} else {
			nextTxNo++
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	// ================================================================================
	// Shut down.
	// ================================================================================

	// Stop the system.
	if args.Verbose {
		fmt.Println("Stopping SMR system.")
	}
	trantorSystem.Stop()

	// Stop the node.
	if args.Verbose {
		fmt.Println("Stopping node.")
	}
	node.Stop()

	return <-nodeErr
}

// Parses the command-line arguments and returns them in a params struct.
func parseArgs(args []string) *parsedArgs {
	app := kingpin.New("chat-demo", "Small chat application to demonstrate the usage of the Mir library.")
	verbose := app.Flag("verbose", "Verbose mode.").Short('v').Bool()
	trace := app.Flag("trace", "Very verbose mode.").Bool()
	// Currently, the type of the node ID is defined as uint64 by the /pkg/types package.
	// In case that changes, this line will need to be updated.
	ownID := app.Arg("id", "ID of this node").Required().String()
	initMembershipFile := app.Flag("init-membership", "File containing the initial system membership.").
		Short('i').Required().String()
	initChkpFile := app.Flag("checkpoint", "Initial state checkpoint to start from.").Short('c').String()
	chkpDir := app.Flag("checkpoint-dir", "Directory to store checkpoints in.").Short('d').String()

	if _, err := app.Parse(args[1:]); err != nil { // Skip args[0], which is the name of the program, not an argument.
		app.FatalUsage("could not parse arguments: %v\n", err)
	}

	return &parsedArgs{
		OwnID:              t.NodeID(*ownID),
		Verbose:            *verbose,
		Trace:              *trace,
		InitMembershipFile: *initMembershipFile,
		InitChkpFile:       *initChkpFile,
		ChkpDir:            *chkpDir,
	}
}

func getPortStr(addressStr string) (string, error) {
	address, err := multiaddr.NewMultiaddr(addressStr)
	if err != nil {
		return "", err
	}

	_, addrStr, err := manet.DialArgs(address)
	if err != nil {
		return "", err
	}

	_, portStr, err := net.SplitHostPort(addrStr)
	if err != nil {
		return "", err
	}

	return portStr, nil
}

func loadStableCheckpoint(filename string) (retChkp *checkpoint.StableCheckpoint, retErr error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open checkpoint file file: %s", filename)
	}
	defer func() {
		retErr = file.Close()
	}()

	chkpBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var chkp *checkpoint.StableCheckpoint
	if err := chkp.Deserialize(chkpBytes); err != nil {
		return nil, err
	}

	return chkp, nil
}
