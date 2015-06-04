$(function() {
  
  var now = Date.now();
  $('#addUser').attr('href', '/user/' + now);
  
  // update link
  $('#addUser').click(function() {
    now = Date.now();
    $(this).attr('href', '/user/' + now);
  });
  
  if (window["WebSocket"]) {
    var wsHost = "ws://" + window.location.host + "/ws/" + uid;
    conn = new WebSocket(wsHost);
    conn.onclose = function(evt) {
      console.log("Socket connection closed");
    }
    conn.onmessage = function(evt) {
      UpdateUserTable(jQuery.parseJSON(evt.data));
    }
  }
  else {
    console.log("Your browser does not support WebSockets");
  };
  
  function UpdateUserTable(users) {
    $userTableRows = $('#users>tbody');
    $userTableRows.empty();
    var html = "";
    for(var i=0;i<users.length;i++) {
      html += "<tr>";
      html += "<td>" + users[i] + "</td>";
      html += "<td>";
      html += "<button type='button' class='btn btn-danger btn-xs'>Send APN</button>\n";
      html += "<button type='button' class='btn btn-success btn-xs'>Send GCM</button>\n";
      html += "<button type='button' class='btn btn-warning btn-xs'>Send Both</button>";
      html += "</td>";
      html += "</tr>";
    }
    $userTableRows.append(html)
  }
});