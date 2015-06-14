package apn

import (
	"encoding/json"
	"github.com/alexjlockwood/gcm"
	"log"
)

// GCMPayload structure holds data required to send
// a Google Cloud Messaging notification
type gcmPayload struct {
	RegistrationIds []string `json:"registration_ids"`
	Title           string   `json:"title"`
	Body            string   `json:"body"`
	Tag             string   `json:"tag"`
}

type gcmService struct {
	apiKey string
}

// Initialize GCM service by passing API Key
// obtained from google console as parameter
func Initialize(apiKey string) (service *gcmService) {
	return &gcmService{apiKey}
}

// Name returns name of the service
func (service *gcmService) Name() string {
	return "gcm"
}

// TryPush will push the incoming payload data
func (service *gcmService) TryPush(data []byte) error {
	payloadData := &gcmPayload{}
	err := json.Unmarshal(data, payloadData)
	if err != nil {
		return err
	}
	log.Println(*payloadData)

	go func(_payloadData *gcmPayload) {

		log.Println("Sending GCM", _payloadData.Title, "to", _payloadData.RegistrationIds)

		// Create the message to be sent
		dataToSend := map[string]interface{}{"title": _payloadData.Title, "body": _payloadData.Body, "tag": _payloadData.Tag}
		msg := gcm.NewMessage(dataToSend, _payloadData.RegistrationIds...)

		// Create a Sender to send the message.
		sender := &gcm.Sender{ApiKey: service.apiKey}

		// Send the message and receive the response
		response, err := sender.SendNoRetry(msg)
		if err != nil {
			log.Println("Failed to send message:", err)
			return
		}
		log.Println(response)

	}(payloadData)

	return nil
}
