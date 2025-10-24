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

func Get[T any](hostname, url, username, password string, port int, useSsl bool, insecure bool) (responseData T, err error) {
	return request[T](hostname, url, "GET", "application/octet-stream", "", username, password, port, nil, useSsl, insecure)
}

func Delete[T any](hostname, url, username, password string, port int, useSsl bool, insecure bool) (responseData T, err error) {
	return request[T](hostname, url, "DELETE", "application/octet-stream", "", username, password, port, nil, useSsl, insecure)
}

func Post[T any](hostname, url, md5, username, password string, port int, data []byte, useSsl bool, insecure bool) (responseData T, err error) {
	return request[T](hostname, url, "POST", "application/json", md5, username, password, port, data, useSsl, insecure)
}

func Patch[T any](hostname, url, username, password string, port int, useSsl bool, insecure bool) (responseData T, err error) {
	return request[T](hostname, url, "PATCH", "application/json", "", username, password, port, nil, useSsl, insecure)
}

func request[T any](hostname, url, method, contentType, md5, username, password string, port int, data []byte, useSsl, insecure bool) (responseData T, err error) {

	uri := FormatSslString(useSsl, fmt.Sprintf("%s:%d%s", hostname, port, url))

	req, err := http.NewRequest(method, uri, bytes.NewReader(data))
	if err != nil {
		return
	}

	req.SetBasicAuth(username, password)

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
		switch any(responseData).(type) {
		case []byte:
			return any(body).(T), nil
		default:
			if err := json.Unmarshal(body, &responseData); err != nil {
				return responseData, err
			}
			return responseData, nil
		}
	}
	return responseData, errors.New(fmt.Sprintf("error in the http request: %v", err))
}
