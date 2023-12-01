import React, { useState } from 'react';
import WebSocketConsole from './components/WebSocketConsole/WebSocketConsole';
import './App.css';

function App() {
    const [webSocketConsoles, setWebSocketConsoles] = useState({}); // dict of { port: WebSocketConsole }
    const [port, setPort] = useState('');

    const connectWebSocket = () => {
        // Implement WebSocket connection logic here
        if (port === '' || port in webSocketConsoles) {
            alert('Port already connected or empty');
            return;
        }

        const newWebSocketConsoles = { ...webSocketConsoles };
        newWebSocketConsoles[port] = <WebSocketConsole key={port} port={port} />;
        setWebSocketConsoles(newWebSocketConsoles);
        setPort('');
    };


    return (
        <div>
            <title>
                WebSocket Demo
            </title>
            <div>
                Enter Port Number: <input type="text" value={port} onChange={(e) => setPort(e.target.value)} />
                <button onClick={connectWebSocket}>Connect</button>
            </div>

            <div className="message-containers" id="message-containers">
                {Object.values(webSocketConsoles)}
            </div>
        </div>
    );
};

export default App;
