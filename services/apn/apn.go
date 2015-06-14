package apn

import (
	"encoding/json"
	apns "github.com/anachronistic/apns"
	"log"
)

const (
	PRODUCTION_GATEWAY = "gateway.push.apple.com:2195"
	SANDBOX_GATEWAY    = "gateway.sandbox.push.apple.com:2195"
)

type apnService struct {
	*apns.Client
}

// apnPayload structure holds data required
// to send an Apple Push Notification
type apnPayload struct {
	Tokens []string `json:"tokens"`
	Alert  string   `json:"alert"`
	Badge  int      `json:"badge"`
	Sound  string   `json:"sound"`
}

// Initialize APN Service by passing gateway url,
// certificate file path and key file path as parameters
func Initialize(gateway, certificate, key string) (service *apnService) {
	apnsClient := apns.NewClient(gateway, certificate, key)
	return &apnService{apnsClient}
}

// Name returns name of the service
func (apn *apnService) Name() string {
	return "apn"
}

// TryPush will push the incoming payload data
func (service *apnService) TryPush(data []byte) error {
	payloadData := &apnPayload{}
	err := json.Unmarshal(data, payloadData)
	if err != nil {
		return err
	}
	log.Println(*payloadData)

	go func(_payloadData *apnPayload) {
		// TODO: Check if apn client manages synchronization
		log.Println("Sending APN", _payloadData.Alert, "to", _payloadData.Tokens)
		payload := apns.NewPayload()
		payload.Alert = _payloadData.Alert
		payload.Badge = _payloadData.Badge
		payload.Sound = _payloadData.Sound

		pn := apns.NewPushNotification()
		pn.AddPayload(payload)
		for _, token := range _payloadData.Tokens {
			pn.DeviceToken = token
			resp := service.Send(pn)
			log.Println("Success:", resp.Success)
			log.Println("  Error:", resp.Error)
		}

	}(payloadData)

	return nil
}
