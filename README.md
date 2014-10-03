go-statsdurl
============

Connect to a database using a `STATSD_URL`.

Usage
=====

It uses [Statsd][statsd] under the hood:

```go
import "github.com/josegonzalez/go-statsdurl"

// Connect using os.Getenv("STATSD_URL").
c, err := statsdurl.Connect()

// Alternatively, connect using a custom Database URL.
c, err := statsdurl.ConnectToURL("udp://...")
```

In both cases you will get the result values of calling
`statsd.NewStatsdClient(...)`, that is, an instance of `statsd.StatsdClient` and an
error.

[statsd]: https://github.com/quipo/statsd

Installation
============

Install it using the "go get" command:

    go get github.com/josegonzalez/go-statsdurl

