package clients

import (
	"net/http"
	"fmt"
)

func makeHttpRequest(url string) (*http.Response, error) {

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode)

	return resp, nil
}
