package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	runTests := m.Run()
	os.Exit(runTests)
}

func Benchmark_makeRequests(b *testing.B) {

	requestCount := 0

	responseTime := make(chan float64, 1)
	err := make(chan error, 1)

	b.ResetTimer()
	for k := 0; k < b.N; k++ {
		go makeRequests(responseTime, err)

		rt := <-responseTime
		er := <-err

		fmt.Println(
			"https://gitlab.com request response time = ", rt)
		fmt.Print(er)
		fmt.Println()

		requestCount++
	}
}
