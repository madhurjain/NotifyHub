package main

import (
	"encoding/json"
	"log"
)

// GCM - structure holds data required to send
// a Google Cloud Messaging notification
type GCM struct {
	RegistrationId string `json:"registration_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	Tag            string `json:"tag"`
}

// SendGCM - sends a Google Cloud Messaging
// notification to a mobile device
func SendGCM(gcm *GCM) {
	// This mocks sending a GCM
	// In practice, a request will be sent to
	// GCM servers to deliver this notification
	go func(_gcm *GCM) {
		log.Println("Sending GCM", _gcm.Title, "to", _gcm.RegistrationId)
		data, err := json.Marshal(_gcm)
		if err != nil {
			log.Println(err)
			return
		}
		msg := &message{uid: _gcm.RegistrationId, payload: data}
		SendMessage(msg)
	}(gcm)
}
