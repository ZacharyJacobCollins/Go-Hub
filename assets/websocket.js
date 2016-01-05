$(function() {

var conn;
var msg = $("#msg");
var chat = $("#chat");

function appendChat(msg) {
    var d = chat[0]
    var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
    msg.appendTo(chat)
    if (doScroll) {
        d.scrollTop = d.scrollHeight - d.clientHeight;
    }
}

//Submit message to socket unless there is no connection or message.
$("#form").submit(function() {
    if (!conn) {
        return false;
    }
    if (!msg.val()) {
        return false;
    }
    conn.send(msg.val());
    msg.val("");
    return false
});

//If there is a socket
if (window["WebSocket"]) {
    conn = new WebSocket("ws://{{$}}/ws");
    conn.onclose = function(evt) {
        appendChat($("<div><b>Connection closed.</b></div>"))
    }
    conn.onmessage = function(evt) {
        appendChat($("<div/>").text(evt.data))
    }
} else {
    appendChat($("<div><b>Your browser does not support WebSockets.</b></div>"))
}
});
