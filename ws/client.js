let ws = new WebSocket("ws://localhost:8000/ws");

ws.onopen = function(evt) {
    console.log("connection opened.");
    ws.send("hello world");
};

ws.onmessage = function(evt) {
    console.log("received message: " + evt.data);
};

ws.onclose = function(evt) {
    console.log("connection closed.");
};
