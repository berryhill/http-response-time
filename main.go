package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/berryhill/http-response-time/clients"
)

func main() {

	go http.ListenAndServe(":8765", nil)

	fmt.Println("Welcome to the HTTP Response Time App")
	fmt.Println()

	requestCount := 0
	timer := time.NewTimer(time.Minute * 5)

	go func() {
		select {
		case <-timer.C:
			fmt.Println()
			fmt.Println("-----------------------------------------------")
			fmt.Println()
			fmt.Println("Made ", requestCount, " requests in 5 minutes")
			os.Exit(1)
		}
	}()

	responseTime := make(chan float64, 1)
	err := make(chan error, 1)

	for {
		go makeRequests(responseTime, err)

		rt := <- responseTime
		reqErr := <- err

		fmt.Println(
			"https://gitlab.com request response time = ", rt)
		fmt.Print("Error: ", reqErr)
		fmt.Println()

		requestCount++
	}
}

func makeRequests(responseTime chan float64, err chan error) {
	rt, reqErr := clients.GetGitlabRequestResponseTime()
	if reqErr != nil {
		responseTime <- -1
		err <- reqErr
	} else {
		responseTime <- rt
		err <- nil
	}
}
