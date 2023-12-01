import React, { useState, useEffect } from 'react';
import './WebSocketConsole.css';

const CustomConsoleComponent = ({ port }) => {
    // const [ws, setWs] = useState(null);
    let ws = null;
    // const [ws, setWs] = useState(null);
    const [incomingMessage, setIncomingMessage] = useState('');
    const [errors, setErrors] = useState([]);
    const [loggedMessages, setLoggedMessages] = useState([]);
    const [syncConnection, setSyncConnection] = useState(true)


    useEffect(() => {
        console.log('i fire once');
        const connectWebSocket = () => {
            const newWs = new WebSocket(`ws://localhost:${port}/ws`);

            newWs.addEventListener('open', () => {
                const message = `WebSocket connection for Port ${port} established.`
                setLoggedMessages(prevMessages => [...prevMessages, message]);
                console.log(message);

                // Clear error messages

                setErrors([]);
            });

            newWs.addEventListener('message', (event) => {
                const message = decomposeJSON(JSON.parse(event.data))
                console.log(syncConnection)
                if( syncConnection === true ){
                    // Sync mode
                    setIncomingMessage(message);
                }
                else {
                    // Async mode
                    setLoggedMessages(prevMessages => [...prevMessages, message]);
                }
                // Update component state or perform other actions based on the received message
            });

            newWs.addEventListener('close', (event) => {
                const message = `WebSocket connection ended`;
                setErrors(prevErrors => [...prevErrors, message]);
                console.log(`WebSocket connection for Port ${port} closed.`);
            });

            newWs.addEventListener('error', (event) => {
                const message = `WebSocket error. See console for further details`;
                console.error(event.valueOf());
                console.log(errors)
                setErrors(prevErrors => [...prevErrors, message]);
                console.log(errors)
                // newWs.close()
            });

            ws = newWs;
            // setWs(newWs);
        };

        connectWebSocket();

        return () => {
            if (ws) {
                ws.close();
            }
        };
    }, []);

    const decomposeJSON = (message) => {
        return JSON.stringify(message, null, 2)
    }

    const acceptIncomingLog = () => {
        if(incomingMessage === ""){
            return;
        }

        // Accept Log and clear the input
        setLoggedMessages(prevMessages => [...prevMessages, incomingMessage]);
        setIncomingMessage("");

        // Send the response on the WebSocket
        let webSocketResponse = {
            "Type": "accept",
            "Value": ""
        };
        const webSocketResponseJSON = JSON.stringify(webSocketResponse);
        ws.send(webSocketResponseJSON);
    };

    const declineIncomingLog = () => {
        if(incomingMessage === ""){
            return;
        }

        // Decline Log and clear the input
        setIncomingMessage("");

        // Send the response on the WebSocket
        let webSocketResponse = {
            "Type": "decline",
            "Value": ""
        };
        const webSocketResponseJSON = JSON.stringify(webSocketResponse);
        ws.send(webSocketResponseJSON);
    };

    const closeConnection = () => {
        console.log('Closing connection from button');
        ws.close()

        // Send the response on the WebSocket
        let webSocketResponse = {
            "Type": "close",
            "Value": ""
        };
        const webSocketResponseJSON = JSON.stringify(webSocketResponse);
        ws.send(webSocketResponseJSON);
    }

    const changeSyncronization = () => {
        if (syncConnection) {
            // Changing to Async mode
            acceptIncomingLog(); // Accept also the existing incoming log
        }
        setSyncConnection(!syncConnection)
    }

    return (
        <div id="sub-screen" className="sub-screen">
            <label className="switch">
                <input type="checkbox" id="sync-switch-input" checked={syncConnection} onChange={changeSyncronization} />
                <div className="slider round">
                    <span className="on">Sync</span>
                    <span className="off">Async</span>
                </div>
            </label>
            <button onClick={closeConnection} className="close-button">Close Connection</button>
            <p className="websocket-console-title">
                WebSocket Console for Port {port}:
            </p>
            <div className="error-screen">
                {errors.map((error, index) => (
                    <p key={index}>{error}</p>
                ))}
            </div>

            {syncConnection && (
                <div id="incoming-log-div" className="incoming-log-div">
                    Incoming Log:
                    <div id="incoming-log-placeholder" className="incoming-log-screen json-message">
                        <pre>{incomingMessage}</pre>
                    </div>
                    <div id="incoming-log-buttons" className="incoming-log-buttons">
                        <button onClick={acceptIncomingLog}>Accept</button>
                        <button onClick={declineIncomingLog}>Decline</button>
                    </div>
                </div>
            )}

            <div id="logged-messages-div" className="logged-messages-div">
                Accepted Logs:
                <div id="log-placeholder" className="log-screen json-message">
                    {loggedMessages.map((log, index) => (
                        <p key={index}>{log}</p>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default CustomConsoleComponent;