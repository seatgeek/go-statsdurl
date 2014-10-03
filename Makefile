test:
	DATABASE_URL=udp://localhost:8125                go test
	DATABASE_URL=udp://localhost:8125/statsd.prefix  go test
