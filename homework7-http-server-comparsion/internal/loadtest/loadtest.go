package loadtest

import (
	"net/http"
	"sync"
	"time"
)

type Stats struct {
	Total     int
	Success   int
	Failed    int
	Fastest   time.Duration
	Slowest   time.Duration
	Avg       time.Duration
	RPS       float64
	TotalTime time.Duration
}

func Run(url string, concurrency, total int) Stats {
	var wg sync.WaitGroup
	jobs := make(chan struct{}, total)
	results := make(chan time.Duration, total)

	for i := 0; i < total; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	start := time.Now()
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			client := &http.Client{}
			for range jobs {
				t0 := time.Now()
				resp, err := client.Get(url)
				if err != nil {
					results <- -1
					continue
				}
				resp.Body.Close()
				results <- time.Since(t0)
			}
		}()
	}

	wg.Wait()
	close(results)

	var s Stats
	s.Total = total
	for d := range results {
		if d < 0 {
			s.Failed++
			continue
		}
		s.Success++
		if s.Fastest == 0 || d < s.Fastest {
			s.Fastest = d
		}
		if d > s.Slowest {
			s.Slowest = d
		}
		s.TotalTime += d
	}
	if s.Success > 0 {
		s.Avg = s.TotalTime / time.Duration(s.Success)
	}
	elapsed := time.Since(start)
	s.RPS = float64(s.Success) / elapsed.Seconds()
	return s
}
