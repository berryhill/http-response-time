package main

import (
	"fmt"
	"os"
	"time"

	"github.com/berryhill/http-response-time/clients"
)

func main() {

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

	for {
		responseTime, err := clients.GetGitlabRequestResponseTime()
		if err != nil {
			panic(err)
		}

		fmt.Println(
			"https://gitlab.com request response time = ", responseTime)
		fmt.Println()

		requestCount++
	}
}
