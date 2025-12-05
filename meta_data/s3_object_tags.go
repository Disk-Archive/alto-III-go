package meta_data

import (
	"encoding/json"
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/google/uuid"
	"net/url"
)

func (o *Object) PutObjectTags(bucketId uuid.UUID, key string, tagSet S3TagSet) error {
	params := url.Values{}
	params.Set("bucket_id", bucketId.String())
	params.Set("key", key)

	query := fmt.Sprintf("/api/v1/object/tagging?%s", params.Encode())

	body, err := json.Marshal(tagSet)
	if err != nil {
		return err
	}

	_, err = http.Put[any](o.Hostname, query, o.Credentials.Username, o.Credentials.Password, o.Port, body, o.UseSsl, o.InsecureSslReq)
	return err
}

func (o *Object) GetObjectTags(bucketId uuid.UUID, key string) ([]S3Tag, error) {

	params := url.Values{}
	params.Set("bucket_id", bucketId.String())
	params.Set("key", key)

	query := fmt.Sprintf("/api/v1/object/tagging?" + params.Encode())

	// Expected response: {"tags": []}
	type tagsResponse struct {
		Tags []S3Tag `json:"tags"`
	}

	resp, err := http.Get[*tagsResponse](o.Hostname, query, o.Credentials.Username, o.Credentials.Password, o.Port, o.UseSsl, o.InsecureSslReq)

	if err != nil {
		return nil, err
	}

	return resp.Tags, nil
}

func (o *Object) DeleteObjectTags(bucketId uuid.UUID, key string) error {

	params := url.Values{}
	params.Set("bucket_id", bucketId.String())
	params.Set("key", key)

	query := fmt.Sprintf("/api/v1/object/tagging?" + params.Encode())

	_, err := http.Delete[any](o.Hostname, query, o.Credentials.Username, o.Credentials.Password, o.Port, o.UseSsl, o.InsecureSslReq)

	return err
}
