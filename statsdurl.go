package statsdurl

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"github.com/quipo/statsd"
)

func Connect() (*statsd.StatsdClient, error) {
	return ConnectToURL(os.Getenv("STATSD_URL"))
}

func ConnectToURL(s string) (c *statsd.StatsdClient, err error) {
	statsdUrl, err := url.Parse(s)

	if err != nil {
		return
	}

	prefix := ""

	if len(statsdUrl.Path) > 1 {
		prefix = strings.TrimPrefix(statsdUrl.Path, "/")
		prefix = fmt.Sprintf("/%s", prefix)
	}

	c = statsd.NewStatsdClient(statsdUrl.Host, prefix)
	err = c.CreateSocket()
	return c, err
}
