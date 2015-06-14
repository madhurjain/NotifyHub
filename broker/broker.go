package broker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// Service interface defines an interface to be implemented by each service
// e.g. APN, GCM
type Service interface {
	Name() string
	TryPush(payload []byte) error
}

// Services is list of available services
type Services map[string]Service

// list of available services
var services = Services{}

// AddServices adds services to be used by broker
// If AddServices is called twice with the same name or if a service is nil,
// it panics.
func AddServices(vices ...Service) {
	for _, service := range vices {
		if service == nil {
			panic("broker: AddServices service is nil")
		}
		if _, dup := services[service.Name()]; dup {
			panic("broker: AddServices called twice for service " + service.Name())
		}
		services[service.Name()] = service
	}
}

// Push iterates through all the available services
// and tries to push the payload through
func Push(payload io.ReadCloser) (map[string]interface{}, error) {

	// decode entities sent in the JSON payload
	// e.g. "apn": {...}, "gcm": {...}
	entities := make(map[string]json.RawMessage)
	err := json.NewDecoder(payload).Decode(&entities)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make(map[string]interface{})
	// Call TryPush passing the payload for each entity
	// while storing the result
	for serviceName, payload := range entities {
		service, err := getService(serviceName)
		if err != nil {
			result[serviceName] = map[string]string{"error": err.Error()}
		}
		err = service.TryPush(payload)
		// TODO: Can be further extended to
		// provide service specific responses
		if err != nil {
			result[serviceName] = map[string]string{"error": err.Error()}
		} else {
			result[serviceName] = map[string]string{"success": "sending"}
		}
	}
	return result, nil
}

// getService returns the service if registered
// else returns an error
func getService(name string) (Service, error) {
	if service, ok := services[name]; ok {
		return service, nil
	}
	return nil, fmt.Errorf("no service for %s exists", name)
}

// getRegisteredServices returns a list of all the services currently in use
func getRegisteredServices() Services {
	// For tests
	return services
}

// unregisterAllServices clears all services available
func unregisterAllServices() {
	// For tests
	services = Services{}
}
