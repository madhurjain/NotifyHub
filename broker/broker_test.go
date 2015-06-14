package broker

import (
	dummy "github.com/madhurjain/notifyhub/services/dummy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddServices(t *testing.T) {
	a := assert.New(t)

	service := dummy.Initialize("")
	AddServices(service)
	a.Equal(len(getRegisteredServices()), 1)
	a.Equal(getRegisteredServices()[service.Name()], service)

	service2 := dummy.Initialize("")
	a.Panics(func() { AddServices(service2) }, "notifyhub: AddServices called twice for service dummy")

	unregisterAllServices()
}

func TestGetService(t *testing.T) {
	a := assert.New(t)

	service := dummy.Initialize("")
	AddServices(service)

	s, err := getService(service.Name())
	a.NoError(err)
	a.Equal(s, service)

	s, err = getService("unknown")
	a.Error(err)
	a.Equal(err.Error(), "no service for unknown exists")

	unregisterAllServices()
}
