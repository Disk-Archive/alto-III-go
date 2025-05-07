package http

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Get[T any](url string) (responseData T, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s%s", "192.168.0.98", url), nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		_ = json.Unmarshal(body, &responseData)
		return responseData, nil
	}

	return responseData, errors.New("agggghhh panic")

}
