package clients

import (
	"net/http"
)

func makeHttpRequest(url string) (*http.Response, error) {

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultTransport.RoundTrip(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
