package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/mir/pkg/events"
	"github.com/filecoin-project/mir/pkg/pb/eventpb"
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
	conn        *websocket.Conn
	upgrader    websocket.Upgrader
	eventSignal chan map[string]interface{}
}

func (wsw *WSWriter) Flush() error {
	//fmt.Println("Starting Flush")
	if wsw.conn == nil {
		return nil
		//panic("Websocket connection is null")
	}
	//fmt.Println("Finishing Flush")
	return nil
}

func (wsw *WSWriter) Close() error {
	// TODO i dont think this connection is ever closed
	//fmt.Println("Starting Close")
	if wsw.conn == nil {
		return nil
	}
	//fmt.Println("Finishing Close")
	return wsw.conn.Close()
}

func (wsw *WSWriter) Write(record eventlog.EventRecord) (eventlog.EventRecord, error) {
	//fmt.Println("Starting Write")
	if wsw.conn == nil {
		fmt.Println("No connection.")
		return record, nil
	}
	if record.Events == nil {
		fmt.Println("No events to print.")
		return record, nil
	}

	acceptedEvents := events.EmptyList()
	iter := record.Events.Iterator()

	for event := iter.Next(); event != nil; event = iter.Next() {
		// Create a new JSON object with a timestamp field
		timestamp := time.Now()
		logData := map[string]interface{}{
			"event":     event,
			"timestamp": timestamp,
		}

		// Marshal the JSON data
		message, err := json.Marshal(logData)
		if err != nil {
			panic(err)
		}

		// Send the JSON message over WebSocket
		if err := wsw.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			return record, fmt.Errorf("error sending message over WebSocket: %w", err)
		}
		//fmt.Println(logData)

		eventAction := <-wsw.eventSignal
		actionType, ok := eventAction["type"].(string)
		value, ok := eventAction["value"].(string)
		if !ok {
			continue
		}

		fmt.Print(value)

		// Decide to keep or drop the event based on the type
		if actionType == "accept" {
			// Process the event if the signal is to keep it
			acceptedEvents.PushBack(event)
		} else if actionType == "decline" {
			// do nothing
		} else if actionType == "start" {

		} else if actionType == "delay" {
			num, err := strconv.Atoi(value)
			dur := time.Duration(num)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			} else {
				time.Sleep(dur * time.Millisecond)
			}
			acceptedEvents.PushBack(event)
		} else if actionType == "modify" {

		} else if actionType == "mode" {

		} else if actionType == "filter" {

		} else if actionType == "close" {

		}
		//fmt.Println(logData) // for now just print, TODO remove this line
	}
	//time.Sleep(5 * time.Second)
	//fmt.Println("Finished Write. Finished all events")
	return eventlog.EventRecord{Events: acceptedEvents}, nil
}

func (wsw *WSWriter) HandleClientSignal(signal map[string]interface{}) {
	// Handle the signal as needed
	// For example, you can send it over the channel as a JSON string
	//signalJSON, _ := json.Marshal(signal)
	wsw.eventSignal <- signal
}

func newWSWriter(port string) *WSWriter {
	fmt.Println("Starting newWSWriter")

	// Create a new WSWriter object
	wsWriter := &WSWriter{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  READ_BUFFER_SIZE,
			WriteBufferSize: WRITE_BUFFER_SIZE,
		},
		eventSignal: make(chan map[string]interface{}),
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
			// go routine for incoming messages
			go func() {
				defer func(conn *websocket.Conn) {
					err := conn.Close()
					if err != nil {

					}
				}(conn) // Ensure the connection is closed when the function exits

				for {
					_, message, err := conn.ReadMessage()
					if err != nil {
						// Handle the error as appropriate
						break
					}

					var signal map[string]interface{}
					err = json.Unmarshal(message, &signal)
					if err != nil {
						// Handle JSON parsing error
						continue
					}
					wsWriter.HandleClientSignal(signal)
				}
			}()
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
		eventlog.EventFilterOpt(func(event *eventpb.Event) bool { //Event Filter is just an example
			switch event.Type.(type) {
			case *eventpb.Event_Transport:
				return true
			default:
				return false
			}
		}),
		eventlog.SyncWriteOpt(),
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
	time.Sleep(120 * time.Second)

	// Stop the node.
	node.Stop()
	transport.Stop()
	fmt.Printf("Mir node stopped: %v\n", <-nodeError)
}
