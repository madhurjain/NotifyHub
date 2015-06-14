package dummy

import (
	"encoding/json"
	"log"
)

// dummyPayload structure
type dummyPayload struct {
	DummyIds []string `json:"dummy_ids"`
	Data     string   `json:"data"`
}

type dummyService struct {
	apiKey string
}

// Initialize Dummy service by passing API Key
func Initialize(apiKey string) (service *dummyService) {
	return &dummyService{apiKey}
}

// Name returns name of the service
func (service *dummyService) Name() string {
	return "dummy"
}

// TryPush will push the incoming payload data
func (service *dummyService) TryPush(data []byte) error {
	payloadData := &dummyPayload{}
	log.Println(string(data))
	err := json.Unmarshal(data, payloadData)
	if err != nil {
		return err
	}
	log.Println(*payloadData)

	go func(_payloadData *dummyPayload) {
		log.Println("Sending Dummy Notification", _payloadData.Data, "to", _payloadData.DummyIds)
	}(payloadData)

	return nil
}
