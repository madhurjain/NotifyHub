package dummy

import (
	"github.com/madhurjain/notifyhub/broker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	service := Initialize("dummy_api_key")
	a.Equal(service.apiKey, "dummy_api_key")
	a.Implements((*broker.Service)(nil), service)
}

func TestTryPush(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	service := Initialize("dummy_api_key")
	err := service.TryPush([]byte(`{"dummy_ids": ["1", "2", "3", "4"], "data": "test"}`))
	a.NoError(err)
}
