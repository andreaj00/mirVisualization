package common

import (
	"time"

	trantorpbtypes "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	tt "github.com/filecoin-project/mir/pkg/trantor/types"
	t "github.com/filecoin-project/mir/pkg/types"
	"github.com/filecoin-project/mir/pkg/util/indexedlist"
)

// ModuleConfig sets the module ids. All replicas are expected to use identical module configurations.
type ModuleConfig struct {
	Self   t.ModuleID // id of this module
	Hasher t.ModuleID
	Timer  t.ModuleID
}

// ModuleParams sets the values for the parameters of an instance of the protocol.
// All replicas are expected to use identical module parameters.
type ModuleParams struct {

	// Maximal number of individual transactions in a single batch.
	MaxTransactionsInBatch int `json:",string"`

	// Maximal total combined payload size of all transactions in a batch (in Bytes)
	MaxPayloadInBatch int `json:",string"`

	// Maximal time between receiving a batch request and emitting a batch.
	// On reception of a batch request, the mempool generally waits
	// until it contains enough transactions to fill a batch (by number or by payload size)
	// and only then emits the new batch.
	// If no batch has been filled by BatchTimeout, the mempool emits a non-full (even a completely empty) batch.
	BatchTimeout time.Duration `json:",string"`

	// If this parameter is not nil, the mempool will not receive transactions directly (through NewTransactions) events.
	// On reception of such an event, it will report an error (making the system crash).
	// Instead, TxFetcher will be called to pull transactions from an external source
	// when they are needed to form a batch (upon the RequestBatch event).
	// Looking up transactions by ID will also always fail (return no transactions).
	TxFetcher func() []*trantorpbtypes.Transaction `json:"-"`
}

// State represents the common state accessible to all parts of the module implementation.
// TODO: Consider moving this definition inside the `internal` subdirectory, as it is only used by the mempool.
type State struct {
	// All the transactions in the mempool.
	// Incoming transactions that have not yet been delivered are added to this list.
	// They are removed upon delivery.
	Transactions *indexedlist.IndexedList[tt.TxID, *trantorpbtypes.Transaction]
}
