package statsdurl_test

import (
	"github.com/josegonzalez/go-statsdurl"
	"github.com/quipo/statsd"
	"testing"
)

func TestConnect(t *testing.T) {
	c, err := statsdurl.Connect()

	if err != nil {
		t.Errorf("Error returned")
	}
}
