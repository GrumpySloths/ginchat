// Create WebSocket connection.
const WebSocket = require('ws');

const socket = new WebSocket(`ws://localhost:8080/ws?name=test`);

const clientID = 1;
// Connection opened
socket.addEventListener("open", function (event) {
    sendText("hello server")
});

//接收信道
socket.onmessage = (event) => {
    console.log(typeof event.data)
    const msg = JSON.parse(event.data);
    console.log("Message type from server ", msg);
}

// socket 发送json信息测试
// Send text to all users through the server
function sendText(msg_send) {
    console.log("开始发送json信息")
    // Construct a msg object containing the data the server needs to process the message from the chat client.
    const msg = {
        Type: "message",
        Text: msg_send,
        Id: clientID,
        Date: Date.now(),
    };

    // Send the msg object as a JSON-formatted string.
    socket.send(JSON.stringify(msg));
}

socket.onopen = (event) => {
    console.log("开始发送json信息");
    sendText("json 消息测试");
};
