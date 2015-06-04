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
      var randomMessage = quotes[Math.floor((Math.random() * 10))];
      notify['apn'] = {"token": toUser, "alert": randomMessage, "badge": 1, "sound": "bingbong.aiff"};
    }
    if (action.indexOf('gcm') != -1) {
      var randomMessage = quotes[Math.floor((Math.random() * 10))];
      notify['gcm'] = {"registration_id": toUser, "title": randomMessage, "body": "description of the message", "tag": "unique"};
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
  
  
  var quotes = [
    "I would like to die on Mars. Just not on impact. - Elon Musk",
    "Life is about making an impact, not making an income.",
    "Whatever the mind of man can conceive and believe, it can achieve.",
    "Only the wisest and stupidest of men never change.",
    "Our greatest glory is not in never falling, but in rising every time we fall.",
    "Success is a lousy teacher. It seduces smart people into thinking they can't lose.",
    "Intellectual property has the shelf life of a banana.",
    "The best time to plant a tree was 20 years ago. The second best time is now.",
    "Any product that needs a manual to work is broken.",
    "I hear and I forget. I see and I remember. I do and I understand.",
    "Choose a job you love, and you will never have to work a day in your life."
  ];
  
});