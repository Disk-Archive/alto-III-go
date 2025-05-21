package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Get[T any](hostname, url string) (responseData T, err error) {
	return request[T](hostname, url, "GET", "application/octet-stream", "", nil, true)
}

func Delete[T any](hostname, url string) (responseData T, err error) {
	return request[T](hostname, url, "DELETE", "application/octet-stream", "", nil, true)
}

func Post[T any](hostname, url, md5 string, data []byte) (responseData T, err error) {
	return request[T](hostname, url, "POST", "application/json", md5, data, true)
}

func request[T any](hostname, url, method, contentType, md5 string, data []byte, insecure bool) (responseData T, err error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://%s%s", hostname, url), bytes.NewReader(data))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("md5hash", md5)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
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

	// ToDo: Need to handle more http statuses
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		_ = json.Unmarshal(body, &responseData)
		return responseData, nil
	}
	return responseData, errors.New("agggghhh panic")
}
