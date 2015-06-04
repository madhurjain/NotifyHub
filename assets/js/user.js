$(function() {
  if (window["WebSocket"]) {
  
    $iosDrawer = $('#ios>.drawer');
    $androidDrawer = $('#android>.drawer');
    
    var wsHost = "ws://" + window.location.host + "/ws/" + uid;
    conn = new WebSocket(wsHost);
    conn.onclose = function(evt) {
      console.log("Socket connection closed");
    }
    conn.onmessage = function(evt) {
      json = jQuery.parseJSON(evt.data);
      // APN sent if token is set
      if(json['token']) {
        $('p', $iosDrawer).text(json['alert']);
        $iosDrawer.slideDown();
        setTimeout(function() { $iosDrawer.slideUp() }, 4000);
      }
      // GCM sent if registration_id is set
      if(json['registration_id']) {      
        $('p', $androidDrawer).text(json['title']);
        $androidDrawer.slideDown();
        setTimeout(function() { $androidDrawer.slideUp() }, 4000);
      }
      console.log(evt.data);
    }
  }
  else {
    console.log("Your browser does not support WebSockets");
  };
});