## NotifyHub

Standalone web service which accepts notifications as JSON and routes to APN or GCM

`POST` JSON in the below format to `/notify`

```json
{
  "apn": {
    "token": "002244aaccddeeff",
    "alert": "title of the notification",
    "badge": 1,
    "sound": "bingbong.aiff"
  },
  "gcm": {
    "registration_id": "aaaabbbbccccddddeeeeffff001122",
    "title": "testing gcm",
    "body": "description of the message",
    "tag": "unique"	
  }  
}
```

#### *[DEMO]*

Implements WebSocket and sends a mock APN / GCM notification

http://notifyhub.herokuapp.com

- Open the above demo link
- Click `Spawn User` to spawn mock users in different tabs
- You can also spawn users in different browser or on a different computer by visiting link `/user/blah`
- Click `Send APN`, `Send GCM` or `Send Both` to send notification to the user