## NotifyHub

Standalone web service which accepts notifications as JSON and routes to APN or GCM Push Notification Services.

More notification services can be added by implementing the broker's `Service` interface.

#### Sample

`POST` below JSON to `/notify` endpoint

```json
{
  "apn": {
    "tokens": ["0a24de", "0a24de"],
    "alert": "testing an apn",
    "badge": 1,
    "sound": "bingbong.aiff"
  },
  "gcm": {
    "registration_ids": ["ab1d2c", "fe45d9"],
    "title": "testing gcm",
    "body": "message body",
    "tag": "popup"
  }
}
```

