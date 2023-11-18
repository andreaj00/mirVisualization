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
            this.retryInterval = 1000; // 1 second, interval to retry to connect // TODO, implement an exponential retry strategy
            this.maxRetries = 10; // After those many attempt just print unable to connect
            this.retryCount = 0;

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

            // Connect WebSocket
            this.connectWebSocket();

            // Create a delete button
            const deleteButton = document.createElement('button');
            deleteButton.textContent = 'Delete';
            deleteButton.addEventListener('click', () => {
                this.deleteComponent();
            });
            // Append the delete button to the component content
            this.content.appendChild(deleteButton);

            // Append new custom element
            this.attachShadow({ mode: 'open' });
            this.shadowRoot.appendChild(this.content);

        }
        // get port(){
        //     this.getAttribute('port')
        // }

        connectedCallback() {
            this.shadowRoot.id = componentID;
            console.log("connectedCallback", this.shadowRoot.innerHTML);
        }

        deleteComponent() {
            // Remove the component from the DOM
            // this.parentNode.removeChild(this);
            this.retryCount = this.maxRetries;
            this.ws.close();
            this.retryCount = 0;
            this.remove(); // TODO, after removing it seems that it is removed from the DOM, but not unregistered... apperantly is still an unresolved problem https://stackoverflow.com/questions/27058648/how-to-remove-or-unregister-a-registered-custom-element
                            // The problem is that then i cannot re-define it again
        }
        // disconnectedCallback() {
        //     this.retryCount = this.maxRetries;
        //     this.ws.close();
        // }

        connectWebSocket() {
            this.ws = new WebSocket(`ws://localhost:${this.port}/ws`);

            this.ws.addEventListener('open', () => {
                this.writeAdditionalMessagePlaceHolder('WebSocket connection established', 'log-placeholder')
                // Reset retry count on successful connection
                this.retryCount = 0;
                this.writeOnPlaceHolder("", 'error-placeholder'); // Delete any existing error
            });

            this.ws.addEventListener('message', (event) => {
                // console.log(`WebSocket ${this.port}  message received: ${event.data}`);
                const message = JSON.parse(event.data);


                // TODO, add checks on function whether to add or not to the Log
                // const shouldAddMessage = confirm("Do you want to add this message to the screen?");
                // if (shouldAddMessage) {
                //     const logPlaceholder = this.shadowRoot.getElementById('log-placeholder');
                //     const messageElement = document.createElement('pre');
                //     messageElement.textContent = `Received JSON:\n${JSON.stringify(message, null, 2)}`;
                //     logPlaceholder.appendChild(messageElement);
                // }

                // if (this.retryCount > 0) { // Reset previous error messages
                // }
                this.writeAdditionalLog(message, 'log-placeholder')
            });

            this.ws.addEventListener('close', (event) => {
                // If connection not available retry till maxRetries
                let message;
                if (this.retryCount < this.maxRetries) {
                    message = `Unable to connect. Retrying in ${this.retryInterval} milliseconds. Retry count: ${this.retryCount + 1} out of ${this.maxRetries}`;
                    setTimeout(() => {
                        this.retryCount++;
                        this.connectWebSocket(); // Retry after the specified interval
                    }, this.retryInterval);
                } else {
                    message = `Max retry count of ${this.maxRetries} reached. Stopping further retries. Please delete this screen before retrying to connect`
                }

                this.writeOnPlaceHolder(message, 'error-placeholder')
            });

            this.ws.addEventListener('error', (error) => {
                console.error('WebSocket error:', error);
                let message = `WebSocket error: ${error}`
                this.writeOnPlaceHolder(message, 'error-placeholder')
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
            messageElement.textContent = `Received JSON:\n${JSON.stringify(message, null, 2)}`; // TODO, adapt if the message is not pretty JSON
            placeholderElement.appendChild(messageElement);
        }

        makeDraggrable() {
            // TO fix

            // Make each console 'Drag and Drop'-able
            // const messageContainers = document.getElementById('message-containers');
            // const subScreen = document.createElement('div');
            // subScreen.className = 'sub-screen';
            // subScreen.id = `sub-screen-${port}`;
            //
            // subScreen.setAttribute('draggable', true);
            // subScreen.addEventListener('dragstart', (event) => {
            //     event.dataTransfer.setData('text/plain', subScreen.id);
            // });
            //
            // subScreen.addEventListener('dragover', (event) => {
            //     event.preventDefault();
            // });
            //
            // subScreen.addEventListener('drop', (event) => {
            //     event.preventDefault();
            //     const draggedScreenId = event.dataTransfer.getData('text/plain');
            //     const draggedScreen = document.getElementById(draggedScreenId);
            //
            //     if (draggedScreen && draggedScreen !== subScreen) {
            //         const parent = subScreen.parentNode;
            //         const nextSibling = subScreen.nextElementSibling;
            //         parent.insertBefore(draggedScreen, subScreen);
            //         parent.insertBefore(subScreen, nextSibling);
            //     }
            // });
            // messageContainers.appendChild(subScreen);

            // const messageConsole = document.createElement('div');
            // messageConsole.textContent = `WebSocket Console for Port ${port}:`;
            // subScreen.appendChild(messageConsole);
        }

    }

    customElements.define(componentID, CustomConsoleComponent);

    return componentID;
}
