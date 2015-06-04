package main

import (
	"log"
)

// APN - structure holds data required to send
// an Apple Push Notification
type APN struct {
	Token string `json:"token"`
	Alert string `json:"alert"`
	Badge int    `json:"badge"`
	Sound string `json:"sound"`
}

// SendAPN - sends an Apple Push Notification
// to a mobile device
func SendAPN(apn *APN) {
	// This mocks sending an APN
	// In practice, a request will be sent to
	// APN servers to deliver this notification
	go func(_apn *APN) {
		log.Println("Sending APN", _apn.Alert, "to", _apn.Token)
		msg := &message{uid: _apn.Token, payload: []byte(_apn.Alert)}
		SendMessage(msg)
	}(apn)
}
