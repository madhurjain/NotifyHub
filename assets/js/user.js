$(function() {
  if (window["WebSocket"]) {
    var wsHost = "ws://" + window.location.host + "/ws/" + uid;
    conn = new WebSocket(wsHost);
    conn.onclose = function(evt) {
      console.log("Socket connection closed");
    }
    conn.onmessage = function(evt) {
      console.log(evt.data);
    }
  }
  else {
    console.log("Your browser does not support WebSockets");
  };
});