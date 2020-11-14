var url = "ws://" + window.location.host + "/ws";
var ws = new WebSocket(url);
var name = "Guest" + Math.floor(Math.random() * 1000);

var chat = document.getElementById("chat");
var text = document.getElementById("text");

var now = function () {
    var iso = new Date().toISOString();
    return iso.split("T")[1].split(".")[0];
};

ws.onmessage = function (msg) {
    var line =  now() + " " + msg.data + "\n";
    chat.innerText += line;
};


function OnButtonClick() {
    ws.send("<" + name + "> " + text.value);
    text.value = "";
};