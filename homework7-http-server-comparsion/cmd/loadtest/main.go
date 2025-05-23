package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yevropii/go-masters-2025/homework7-http-server-comparsion/internal/loadtest"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "target URL")
	c := flag.Int("c", 50, "concurrency")
	n := flag.Int("n", 1000, "number of requests")
	flag.Parse()

	stats := loadtest.Run(*url, *c, *n)
	if stats.Failed > 0 {
		log.Printf("failed requests: %d", stats.Failed)
	}

	fmt.Printf(`
			URL           : %s
			Total         : %d
			Success       : %d
			Failed        : %d
			Fastest       : %v
			Slowest       : %v
			Average       : %v
			Requests/sec  : %.2f
		`,
		*url,
		stats.Total,
		stats.Success,
		stats.Failed,
		stats.Fastest,
		stats.Slowest,
		stats.Avg,
		stats.RPS,
	)
}
