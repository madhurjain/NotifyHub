package main

import (
	"log"
)

// WS - structure holds data required to send a payload over WebSocket
// UserId identifies the user to which the payload is to be sent
type WS struct {
	UserId  int                    `json:"user_id"`
	Payload map[string]interface{} `json:"payload"`
}

// SendWS - sends the payload over WebSocket
func SendWS(ws *WS) {
	go func(_ws *WS) {
		log.Println("Sending over websocket", _ws.Payload, "to", _ws.UserId)
	}(ws)
}
