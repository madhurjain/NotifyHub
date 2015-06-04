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
  
  $('#users').on('click', 'button', function() {
    var toUser = "" + $(this).data('uid');  // force uid to string
    var action = $(this).data('action');
    var notify = {};
    if (action.indexOf('apn') != -1) {
      notify['apn'] = {"token": toUser, "alert": "title of the notification", "badge": 1, "sound": "bingbong.aiff"};
    }
    if (action.indexOf('gcm') != -1) {
      notify['gcm'] = {"registration_id": toUser, "title": "testing gcm", "body": "description of the message", "tag": "unique"};
    }
    $.ajax({
      url: '/notify',
      type: 'post',
      dataType: 'json',
      contentType : 'application/json',
      data: JSON.stringify(notify),
      success: function(data) {
        console.log(data);
      }
    });
  });
  
  function UpdateUserTable(users) {
    $userTableRows = $('#users>tbody');
    $userTableRows.empty();
    var html = "";
    for(var i=0;i<users.length;i++) {
      html += "<tr>";
      html += "<td>" + users[i] + "</td>";
      html += "<td>";
      html += "<button data-action='apn' data-uid='" + users[i] + "' type='button' class='btn btn-danger btn-xs'>Send APN</button>\n";
      html += "<button data-action='gcm' data-uid='" + users[i] + "' type='button' class='btn btn-success btn-xs'>Send GCM</button>\n";
      html += "<button data-action='apn,gcm' data-uid='" + users[i] + "' type='button' class='btn btn-warning btn-xs'>Send Both</button>";
      html += "</td>";
      html += "</tr>";
    }
    $userTableRows.append(html)
  }
});