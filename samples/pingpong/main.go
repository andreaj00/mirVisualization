package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"net/http"

	"github.com/filecoin-project/mir"
	"github.com/filecoin-project/mir/pkg/eventlog"
	"github.com/filecoin-project/mir/pkg/logging"
	"github.com/filecoin-project/mir/pkg/modules"
	"github.com/filecoin-project/mir/pkg/net/grpc"
	trantorpbtypes "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	"github.com/filecoin-project/mir/pkg/timer"
	t "github.com/filecoin-project/mir/pkg/types"
)

const (
	READ_BUFFER_SIZE  = 1024
	WRITE_BUFFER_SIZE = 1024
	BASE_PORT         = 8080
)

type WSWriter struct {
	// ... websocket server variables ...
	conn     *websocket.Conn
	upgrader websocket.Upgrader
}

func (wsw *WSWriter) Flush() error {
	fmt.Println("Starting Flush")
	//if wsw.conn == nil { // TODO, This creates the process to panic and exit if the eventlog.SyncWriteOpt() option is ON
	//	panic("Websocket connection is null")
	//}
	fmt.Println("Finishing Flush")
	return nil
}

func (wsw *WSWriter) Close() error {
	// TODO i dont think this connection is ever closed
	fmt.Println("Starting Close")
	if wsw.conn == nil {
		return nil
	}
	fmt.Println("Finishing Close")
	return wsw.conn.Close()
}

// TODO, for now just returning the input record, check with Catherine what to return (there are 4 returns)
func (wsw *WSWriter) Write(record eventlog.EventRecord) (eventlog.EventRecord, error) {
	fmt.Println("Starting Write")

	// Wait for the connection to be established
	for wsw.conn == nil {
		fmt.Println("No connection.")
		time.Sleep(time.Millisecond * 100) // TODO, double check if we should use this waiting strategy
	}

	if record.Events == nil {
		fmt.Println("No events to print.")
		return record, nil
	}
	iter := record.Events.Iterator()
	for event := iter.Next(); event != nil; event = iter.Next() {
		// Create a new JSON object with a timestamp field
		timestamp := time.Now()
		logData := map[string]interface{}{
			"event":     event,
			"timestamp": timestamp,
		}

		fmt.Println(logData) // for now just print, TODO remove this line

		// Marshal the JSON data
		message, err := json.Marshal(logData)
		if err != nil {
			panic(err)
		}

		// Send the JSON message over WebSocket
		if err := wsw.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			return record, fmt.Errorf("error sending message over WebSocket: %w", err)
		}

		// Wait for a response from the server
		_, response, err := wsw.conn.ReadMessage()
		if err != nil {
			return record, fmt.Errorf("error reading message from WebSocket: %w", err)
		}
		// Process the response as needed
		fmt.Println("Received response:", string(response))
		// TODO, add response logic here

	}
	fmt.Println("Finished Write. Finished all events")
	return record, nil
}

func newWSWriter(port string) *WSWriter {
	fmt.Println("Starting newWSWriter")

	// Create a new WSWriter object
	wsWriter := &WSWriter{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  READ_BUFFER_SIZE,
			WriteBufferSize: WRITE_BUFFER_SIZE,
		},
	}

	// Create an Async go routine that waits for the connection, TODO check if this decrement to much the performance as now we double every MIR node thread numbers
	go func() {
		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			wsWriter.upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Allow opening the connection by HTML file
			conn, err := wsWriter.upgrader.Upgrade(w, r, nil)
			if err != nil {
				panic(err)
			}

			fmt.Println("WebSocket connection established")

			// Update the attribute of the WSWriter object with the established connection
			wsWriter.conn = conn
		})

		err := http.ListenAndServe(port, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Finished ListenAndServe")

		// Make the program keep running until CTRL+C
		select {}
	}()

	fmt.Println("Finishing newWSWriter")
	return wsWriter
}

func writerFactory(dest string, nodeID t.NodeID, logger logging.Logger) (eventlog.EventWriter, error) {
	fmt.Println("Starting writerFactory")
	fmt.Println("Finishing writerFactory")
	ownPort, err := strconv.Atoi(string(nodeID))
	if err != nil {
		panic(err)
	}
	ownPort += BASE_PORT
	return newWSWriter(fmt.Sprintf(":%d", ownPort)), nil // TODO, for now the port is statically inserted and modified using the nodeID, check if enough
}

func main() {
	fmt.Println("Starting ping-pong")

	// Manually create system membership with just 2 nodes.
	membership := &trantorpbtypes.Membership{map[t.NodeID]*trantorpbtypes.NodeIdentity{ // nolint:govet
		"0": {"0", "/ip4/127.0.0.1/tcp/10000", nil, "1"}, // nolint:govet
		"1": {"1", "/ip4/127.0.0.1/tcp/10001", nil, "1"}, // nolint:govet
	}}

	// Get own ID from command line.
	ownID := t.NodeID(os.Args[1])

	// Instantiate network trnasport module and establish connections.
	transport, err := grpc.NewTransport(ownID, membership.Nodes[ownID].Addr, logging.ConsoleWarnLogger)
	if err != nil {
		panic(err)
	}
	if err := transport.Start(); err != nil {
		panic(err)
	}
	transport.Connect(membership)

	interceptor, err := eventlog.NewRecorder(
		ownID,
		fmt.Sprintf("./node%s", ownID),
		logging.ConsoleInfoLogger,
		eventlog.EventWriterOpt(writerFactory),
		//eventlog.EventFilterOpt(func(event *eventpb.Event) bool { //Event Filter is just an example
		//	switch event.Type.(type) {
		//	case *eventpb.Event_Transport:
		//		return true
		//	default:
		//		return false
		//	}
		//}),
		eventlog.SyncWriteOpt(), // execution synchronous with iteration through events
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Interceptor created successfully")

	// Instantiate Mir node.
	node, err := mir.NewNode(
		ownID,
		mir.DefaultNodeConfig(),
		map[t.ModuleID]modules.Module{
			"transport": transport,
			"pingpong":  NewPingPong(ownID),
			//"pingpong": lowlevel.NewPingPong(ownID),
			"timer": timer.New(),
		},
		interceptor,
	)
	if err != nil {
		panic(err)
	}

	// Run the node for 5 seconds.
	nodeError := make(chan error)
	go func() {
		nodeError <- node.Run(context.Background())
	}()
	fmt.Println("Mir node running.")
	time.Sleep(20 * time.Second)

	// Stop the node.
	node.Stop()
	transport.Stop()
	fmt.Printf("Mir node stopped: %v\n", <-nodeError)
}
