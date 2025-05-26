package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	URL string
	Duration time.Duration
	Error error
}

func simulateFetchURL(url string, results chan<- Result ) {
	start := time.Now()
	delay := time.Duration(rand.Intn(3000) + 1000) * time.Millisecond

	time.Sleep(delay)
	duration := time.Since(start)

	var err error
	if rand.Intn(10) < 3 {
		err = fmt.Errorf("failed to fetch url %s", url)
	}

	results <- Result{URL: url, Duration: duration, Error: err}
}

func main() {
	urls := []string {
		"http://google.com",
		"http://youtube.com",
		"http://gmail.com",
	}

	var wg sync.WaitGroup

	results := make(chan Result, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			simulateFetchURL(u, results)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}
