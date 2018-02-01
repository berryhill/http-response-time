package clients

import (
	"time"
)

const baseUrl = "https://gitlab.com"

func GetGitlabRequestResponseTime() (float64, error) {

	before := time.Now()
	_, err := makeHttpRequest(baseUrl)

	if err != nil {
		return 0.0, err
	}

	after := time.Now()
	duration := after.Sub(before)

	return duration.Seconds(), nil
}
