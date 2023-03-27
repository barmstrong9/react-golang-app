var socket = new WebSocket("ws://localhost:8080/ws")

// Connects us to the websocket endpoint and will listen for events and may error
let connect = () => {
    console.log("Attempting to connect...");

    socket.onopen = () => {
        console.log("Connected successfully");
    };

    socket.onmessage = msg => {
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("Socket closed connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket error: ", error);
    };
};

//Allows us to send messages to the backend via the websocket
let sendMsg = msg => {
    console.log("Sending message...", msg);
    socket.send(msg);
};

export { connect, sendMsg };