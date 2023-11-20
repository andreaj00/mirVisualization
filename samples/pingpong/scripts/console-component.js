// class ConsoleComponent extends HTMLElement {
//     constructor() {
//         super();
//
//         const template = document.getElementById('console-template');
//         const content = document.importNode(template.content, true);
//
//         this.attachShadow({ mode: 'open' });
//         this.shadowRoot.appendChild(content);
//     }
// }
//
// customElements.define('console-component', ConsoleComponent);
function generateComponentID(port) {
    return `console-component-${port}`;
}

function createCustomConsoleComponent(port) {
    const componentID = generateComponentID(port);

    // Note, we expect that the inserted port connection does not exist
    // if (customElements.get(componentID) !== undefined) { // Component for the port already existing
    //     return componentID;
    // }

    // Then able to access it using document.getElementsByTagName('console-component-{port}')
    class CustomConsoleComponent extends HTMLElement {
        constructor() {
            super();
            this.port = port
            // this.id = componentID // It gives that it is not permitted to modify the id

            // Retrieve the template from the shadow DOM of the current document and import it
            const template = document.getElementById('console-template');
            if (!template) {
                console.error('Template not found');
                return;
            }
            this.content = document.importNode(template.content, true);

            // Customise the Component UI port number
            const portIdPlaceholder = this.content.getElementById('port-id-placeholder');
            // Check if the placeholder element exists before updating its content
            if (portIdPlaceholder) {
                portIdPlaceholder.textContent = port;
            }

            // Add the various buttons on the incoming-log-buttons
            const incomingLogButtons = this.content.getElementById('incoming-log-buttons');
            // Check if the placeholder element exists before updating its content
            if (incomingLogButtons) {// Create Accept button
                const acceptButton = document.createElement('button');
                acceptButton.textContent = 'Accept';
                acceptButton.addEventListener('click', () => this.acceptIncomingLog());

                // Create Decline button
                const declineButton = document.createElement('button');
                declineButton.textContent = 'Decline';
                declineButton.addEventListener('click', () => this.declineIncomingLog() );

                // Create Delay section with button and input field
                const delayDiv = document.createElement('div');
                delayDiv.className = 'delay-div';
                const delayButton = document.createElement('button');
                delayButton.textContent = 'Delay';
                delayButton.addEventListener('click', () => this.delayIncomingLog() );
                // Create an input field for milliseconds
                const delayMillisecondsInput = document.createElement('input');
                delayMillisecondsInput.type = 'number';
                delayMillisecondsInput.id = 'delay-milliseconds-input';
                delayMillisecondsInput.placeholder = 'Enter milliseconds';
                delayMillisecondsInput.min = 0;
                delayMillisecondsInput.step = 1;
                // Create a span element for 'ms'
                const msSpan = document.createElement('span');
                msSpan.textContent = 'ms';

                delayDiv.appendChild(delayMillisecondsInput);
                delayDiv.appendChild(msSpan);
                delayDiv.appendChild(delayButton);

                // Append the buttons to the incoming-log-buttons element
                incomingLogButtons.appendChild(acceptButton);
                incomingLogButtons.appendChild(declineButton);
                incomingLogButtons.appendChild(delayDiv);
            }

            // Create a close connection button
            const closeButton = document.createElement('button');
            closeButton.textContent = 'Close Connection';
            closeButton.addEventListener('click', () => this.closeConnection());
            closeButton.classList.add('close-button');
            const subScreen = this.content.getElementById('sub-screen');
            subScreen.insertBefore(closeButton, this.content.getElementById('websocket-title'));

            // Create a switch for Sync/Async mode
            const syncSwitchLabel = document.createElement("label");
            syncSwitchLabel.className = "switch";

            const syncSwitchInput = document.createElement("input");
            syncSwitchInput.type = "checkbox";
            syncSwitchInput.id = "sync-switch-input";
            syncSwitchInput.checked = true;

            const syncSwitchSliderDiv = document.createElement("div");
            syncSwitchSliderDiv.className = "slider round";

            // Create span element for "Confirmed" text
            const syncSwitchOnSpan = document.createElement("span");
            syncSwitchOnSpan.className = "on";
            syncSwitchOnSpan.appendChild(document.createTextNode("Sync"));

            // Create span element for "NA" text
            const syncSwitchOffSpan = document.createElement("span");
            syncSwitchOffSpan.className = "off";
            syncSwitchOffSpan.appendChild(document.createTextNode("Async"));

            // Append elements to build the HTML structure
            syncSwitchSliderDiv.appendChild(syncSwitchOnSpan);
            syncSwitchSliderDiv.appendChild(syncSwitchOffSpan);

            syncSwitchLabel.appendChild(syncSwitchInput);
            syncSwitchLabel.appendChild(syncSwitchSliderDiv);

            syncSwitchInput.addEventListener('change', () => this.changeSyncronization());

            subScreen.insertBefore(syncSwitchLabel, closeButton);


            // Connect WebSocket
            this.connectWebSocket();

            // Create a delete button
            // const deleteButton = document.createElement('button');
            // deleteButton.textContent = 'Delete';
            // deleteButton.addEventListener('click', () => {
            //     this.deleteComponent();
            // });
            // // Append the delete button to the component content
            // this.content.appendChild(deleteButton);

            // Append new custom element
            this.attachShadow({ mode: 'open' });
            this.shadowRoot.appendChild(this.content);
        }

        // connectedCallback() {
        //     this.shadowRoot.id = componentID;
        //     console.log("connectedCallback", this.shadowRoot.innerHTML);
        // }

        acceptIncomingLog(){
            const incomingLog = this.shadowRoot.getElementById('incoming-log-placeholder').textContent;
            if (incomingLog === "") {
                return
            }

            // Accept Log and clear the input
            this.shadowRoot.getElementById('incoming-log-placeholder').textContent = ""
            this.addLogAndRespondToWS(incomingLog);
        }

        addLogAndRespondToWS(incomingLog) {
            this.writeAdditionalLog(incomingLog, 'log-placeholder')

            // Send the response on the WebSocket
            let webSocketResponse = {
                "Type": "accept",
                "Value": ""
            };
            const webSocketResponseJSON = JSON.stringify(webSocketResponse);
            this.ws.send(webSocketResponseJSON);
        }

        declineIncomingLog(){
            if (this.shadowRoot.getElementById('incoming-log-placeholder').textContent === "") {
                return
            }

            // Decline Log and clear the input
            // Do nothing for the decline part
            this.shadowRoot.getElementById('incoming-log-placeholder').textContent = ""

            // Send the response on the WebSocket
            let webSocketResponse = {
                "Type": "decline",
                "Value": ""
            };
            const webSocketResponseJSON = JSON.stringify(webSocketResponse);
            this.ws.send(webSocketResponseJSON);
        }

        delayIncomingLog(){
            const incomingLog = this.shadowRoot.getElementById('incoming-log-placeholder').textContent;
            if (incomingLog === "") {
                return
            }

            // Get the delay values
            const delayMilliseconds = this.shadowRoot.getElementById('delay-milliseconds-input').value;
            // TODO add input check on delayMilliseconds

            // Accept Log and clear the input
            this.writeAdditionalLog(incomingLog, 'log-placeholder')
            this.shadowRoot.getElementById('incoming-log-placeholder').textContent = ""

            // Send the response on the WebSocket
            let webSocketResponse = {
                "Type": "delay",
                "Value": delayMilliseconds
            };
            const webSocketResponseJSON = JSON.stringify(webSocketResponse);
            this.ws.send(webSocketResponseJSON);
        }

        changeSyncronization(){
            let responseValue = ""
            if(this.shadowRoot.getElementById('sync-switch-input').checked){
                // Sync mode
                responseValue = 'sync'
                // Unhide the incoming-log-div
                this.shadowRoot.getElementById('incoming-log-div').style.display = 'block';
            }
            else {
                // Async mode
                responseValue = 'async'
                // Hide the incoming-log-div
                this.shadowRoot.getElementById('incoming-log-div').style.display = 'none';

                //TODO, check this default behaviour
                // By default, accept the Incoming Log not accepted yet (if empty there is a check inside the function)
                this.acceptIncomingLog()
            }



            // Send the response on the WebSocket
            let webSocketResponse = {
                "Type": "mode",
                "Value": responseValue
            };
            const webSocketResponseJSON = JSON.stringify(webSocketResponse);
            this.ws.send(webSocketResponseJSON);
        }

        closeConnection(){
            console.log('Closing connection from button');
            this.ws.close()

            // Send the response on the WebSocket
            let webSocketResponse = {
                "Type": "close",
                "Value": ""
            };
            const webSocketResponseJSON = JSON.stringify(webSocketResponse);
            this.ws.send(webSocketResponseJSON);
        }

        // deleteComponent() {
        //     // Remove the component from the DOM
        //     // this.parentNode.removeChild(this);
        //     this.ws.close();
        //     this.remove(); // TODO, after removing it seems that it is removed from the DOM, but not unregistered... apperantly is still an unresolved problem https://stackoverflow.com/questions/27058648/how-to-remove-or-unregister-a-registered-custom-element
        //                     // The problem is that then i cannot re-define it again
        // }
        // disconnectedCallback() {
        //     this.ws.close();
        // }

        // Expect that there is already the websocket server waiting
        connectWebSocket() {
            this.ws = new WebSocket(`ws://localhost:${this.port}/ws`);

            this.ws.addEventListener('open', () => {
                this.writeAdditionalMessagePlaceHolder('WebSocket connection established', 'log-placeholder')
                this.writeOnPlaceHolder("", 'error-placeholder'); // Delete any existing error
            });

            this.ws.addEventListener('message', (event) => {
                // console.log(`WebSocket ${this.port}  message received: ${event.data}`);
                // Add the new message to the Incoming Log
                if( this.shadowRoot.getElementById('sync-switch-input').checked){
                    // Sync mode
                    this.shadowRoot.getElementById('incoming-log-placeholder').textContent = this.decomposeJSON(JSON.parse(event.data))
                }
                else {
                    // Async mode
                    this.addLogAndRespondToWS(this.decomposeJSON(JSON.parse(event.data)));
                }
            });

            this.ws.addEventListener('close', (event) => {
                // Does nothing
                this.writeAdditionalMessagePlaceHolder('WebSocket connection ended', 'error-placeholder');
                console.error("Closing connection to port: ", this.port); // TODO, to delete this line
            });

            this.ws.addEventListener('error', (event) => {
                console.error('WebSocket error:', event);
                this.writeOnPlaceHolder('WebSocket error, see console', 'error-placeholder');
                this.ws.close(); // Close the WebSocket in case of an error
            });
        }

        writeOnPlaceHolder(message, placeholder) {
            const placeholderElement = this.shadowRoot.getElementById(placeholder);
            placeholderElement.textContent = message;
        }

        writeAdditionalMessagePlaceHolder(message, placeholder) {
            const placeholderElement = this.shadowRoot.getElementById(placeholder);
            const messageElement = document.createElement('pre');
            messageElement.textContent = message
            placeholderElement.appendChild(messageElement);
        }

        writeAdditionalLog(message, placeholder) {
            const placeholderElement = this.shadowRoot.getElementById(placeholder);
            const messageElement = document.createElement('pre');
            messageElement.textContent = `Received JSON:\n${message}`; // TODO, adapt if the message is not pretty JSON
            placeholderElement.appendChild(messageElement);
        }

        decomposeJSON(message){
            return JSON.stringify(message, null, 2)
        }

    }

    customElements.define(componentID, CustomConsoleComponent);

    return componentID;
}
