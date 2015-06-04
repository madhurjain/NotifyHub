## NotifyHub

`POST` the below JSON to `/notify`

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
  },
  "ws": {
    "user_id": 1,
    "payload": { "title": "ping", "message": "testing websockets" }
  }
}
```
