package alto

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/types"
	"github.com/google/uuid"
	"io"
	"net/url"
)

func (a *AltoIII) CreateS3MultipartUpload(ctx context.Context, objectKey, contentType string, diskId, groupId, bucketId uuid.UUID) (uploadId string, err error) {

	params := url.Values{}
	params.Set("disk_id", diskId.String())
	params.Set("object_key", objectKey)
	params.Set("group_id", groupId.String())
	params.Set("bucket_id", bucketId.String())

	query := "/api/v1/s3-multipart-upload/create?" + params.Encode()

	type respData struct {
		UploadId string `json:"upload_id"`
	}

	resp, err := http.Post[*respData](a.Hostname, query, "", a.Credentials.Username, a.Credentials.Password, contentType, a.Port, nil, a.UseSsl, a.InsecureSslReq)
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString([]byte(resp.UploadId)), nil
}

func (a *AltoIII) UploadPartS3MultipartUpload(ctx context.Context, md5, uploadIdBase64, partNumber string, contentLength int64, reader io.Reader) (eTag string, err error) {

	uploadIdBytes, err := base64.StdEncoding.DecodeString(uploadIdBase64)
	uploadId := string(uploadIdBytes)

	params := url.Values{}
	params.Set("upload_id", uploadId)
	params.Set("part_number", partNumber)

	query := "/api/v1/s3-multipart-upload?" + params.Encode()

	respBody, err := http.UploadStream(ctx, a.Hostname, query, "PUT", "application/octet-stream", md5, a.Credentials.Username, a.Credentials.Password, contentLength, a.Port, reader, a.UseSsl, a.InsecureSslReq)
	if err != nil {
		return "", err
	}

	var resp struct {
		ETag string `json:"e_tag"`
	}

	if err := json.NewDecoder(respBody).Decode(&resp); err != nil {
		return "", err
	}

	return resp.ETag, nil
}

func (a *AltoIII) CompleteS3MultipartUpload(ctx context.Context, uploadIdBase64 string, bucketId uuid.UUID, parts types.CompleteMultipartUploadRequest) (eTag string, err error) {

	uploadIdBytes, err := base64.StdEncoding.DecodeString(uploadIdBase64)
	uploadId := string(uploadIdBytes)

	bodyBytes, err := json.Marshal(parts)
	if err != nil {
		return "", fmt.Errorf("failed to marshal parts: %w", err)
	}

	params := url.Values{}
	params.Set("upload_id", uploadId)

	query := "/api/v1/s3-multipart-upload/complete?" + params.Encode()

	type respData struct {
		ETag string `json:"e_tag"`
	}

	resp, err := http.Post[*respData](a.Hostname, query, "", a.Credentials.Username, a.Credentials.Password, "application/json", a.Port, bodyBytes, a.UseSsl, a.InsecureSslReq)
	if err != nil {
		return
	}
	return resp.ETag, nil
}

func (a *AltoIII) AbortS3MultipartUpload(ctx context.Context, uploadIdBase64 string) (err error) {

	uploadIdBytes, err := base64.StdEncoding.DecodeString(uploadIdBase64)
	uploadId := string(uploadIdBytes)

	params := url.Values{}
	params.Set("upload_id", uploadId)

	query := "/api/v1/s3-multipart-upload?" + params.Encode()

	_, err = http.Delete[interface{}](a.Hostname, query, a.Credentials.Username, a.Credentials.Password, a.Port, a.UseSsl, a.InsecureSslReq)

	return
}
