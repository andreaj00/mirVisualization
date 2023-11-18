

// Create a new WebSocket client connecting to the inputed port
function connectWebSocket() {
    const portInput = document.getElementById('port-input');
    const port = portInput.value;
    document.getElementById('port-input').value = ""; // Clear input value

    // TODO, Do input value check (maybe use regex)
    if (port) {
        // Check existing connections
        // const existingConsole = customElements.get(generateComponentID(port))
        // if ( existingConsole !== undefined && existingConsole != null ) {
        //     console.log("customElements.get(generateComponentID(port))", customElements.get(generateComponentID(port)))
        //     // TODO, in the future highlight and bring the user to the already existing console and/or throw alert
        //     alert(`WebSocket connection for Port ${port} already exists.`);
        //     return;
        // }

        const componentID = createCustomConsoleComponent(port);
        const messageContainers = document.getElementById('message-containers');

        // Create an instance of the dynamically generated custom component
        const customComponent = document.createElement(componentID);

        // Append the custom component to the message containers
        messageContainers.appendChild(customComponent);
        console.log(customComponent)
    }
}
