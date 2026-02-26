package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func UploadStream(ctx context.Context, hostname, url, method, contentType, md5, username, password string, contentLength int64, port int, body io.Reader, useSsl, insecure bool) (io.ReadCloser, error) {

	uri := FormatSslString(useSsl, fmt.Sprintf("%s:%d%s", hostname, port, url))

	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", contentType)

	if md5 != "" {
		req.Header.Set("Content-MD5", md5)
	}

	req.ContentLength = contentLength

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
			IdleConnTimeout: 120 * time.Second,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()

		errBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("http status %d: %s", resp.StatusCode, string(errBody))
	}

	return resp.Body, nil
}

func DownloadStream[T any](ctx context.Context, hostname, url, method, contentType, username, password string, port int, useSsl, insecure bool) (body io.ReadCloser, err error) {
	uri := FormatSslString(useSsl, fmt.Sprintf("%s:%d%s", hostname, port, url))

	req, err := http.NewRequestWithContext(ctx, method, uri, http.NoBody)
	if err != nil {
		return
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: insecure},
			IdleConnTimeout:    time.Second * 60,
			DisableCompression: true,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Printf("failed to close responce body: %v", err)
			}
		}()

		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
