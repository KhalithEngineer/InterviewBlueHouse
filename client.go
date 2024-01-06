package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// findMinMaxTime finds the minimum and maximum time durations from a channel of elapsed times.
func findMinMaxTime(timeElapsedChn chan time.Duration) (time.Duration, time.Duration) {
	var min, max time.Duration

	// Iterate through the channel and update min and max accordingly.
	for t := range timeElapsedChn {
		if t < min || min == 0 {
			min = t
		}
		if t > max {
			max = t
		}
	}
	return min, max
}

// hitRequest sends a POST request to a specified URL and records the elapsed time.
func hitRequest(timeElapsed chan time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	// Prepare request payload and set up the request.
	payload := strings.NewReader("hello")
	req, _ := http.NewRequest("POST", "http://localhost:8080/helloworld", payload)
	client := &http.Client{}
	startTime := time.Now()

	// Set request headers.
	req.Header.Add("Content-Type", "text/plain")

	// Send the request.
	res, err := client.Do(req)
	if err != nil {
		// Handle error TODO
		fmt.Println("Error in sending request")
	}

	// Ensure response body is closed.
	defer res.Body.Close()

	// Read and discard the response body.
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Calculate and print the time taken for the request.
	tt := time.Duration(time.Since(startTime).Seconds())
	fmt.Println("Time taken for the request:", tt)
	timeElapsed <- tt
}

func main() {
	maxReq := 10
	timeElapsedChn := make(chan time.Duration, maxReq)
	wg := sync.WaitGroup{}

	// Launch goroutines to send requests concurrently.
	for i := 1; i <= maxReq; i++ {
		wg.Add(1)
		go hitRequest(timeElapsedChn, &wg)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Close the timeElapsedChn channel to signal that no more values will be sent.
	close(timeElapsedChn)

	// Find and print the minimum and maximum time taken for requests.
	min, max := findMinMaxTime(timeElapsedChn)
	fmt.Println("Minimum time taken for request:", min)
	fmt.Println("Maximum time taken for request:", max)
}
