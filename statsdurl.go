package statsdurl

import (
	"net/url"
	"os"

	"github.com/quipo/statsd"
)

func Connect(prefix string) (*statsd.StatsdClient, error) {
	return ConnectToURL(os.Getenv("STATSD_URL"))
}

func ConnectToURL(s string, prefix string) (c *statsd.StatsdClient, err error) {
	statsdUrl, err := url.Parse(s)

	if err != nil {
		return statsdUrl, err
	}

	c = statsd.NewStatsdClient(statsdUrl.Host, prefix)
	err = c.CreateSocket()
	return c, err
}
