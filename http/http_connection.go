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

func Get[T any](hostname, url string, useSsl bool) (responseData T, err error) {
	return request[T](hostname, url, "GET", "application/octet-stream", "", nil, true, useSsl)
}

func Delete[T any](hostname, url string, useSsl bool) (responseData T, err error) {
	return request[T](hostname, url, "DELETE", "application/octet-stream", "", nil, true, useSsl)
}

func Post[T any](hostname, url, md5 string, data []byte, useSsl bool) (responseData T, err error) {
	return request[T](hostname, url, "POST", "application/json", md5, data, true, useSsl)
}

func request[T any](hostname, url, method, contentType, md5 string, data []byte, useSsl, insecure bool) (responseData T, err error) {
	var uri string
	if useSsl {
		uri = fmt.Sprintf("https://%s%s", hostname, url)
	} else {
		uri = fmt.Sprintf("http://%s%s", hostname, url)
	}

	req, err := http.NewRequest(method, uri, bytes.NewReader(data))
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
