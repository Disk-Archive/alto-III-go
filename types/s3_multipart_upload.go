package types

type CompletedPart struct {
	PartNumber int    `xml:"PartNumber" json:"part_number"`
	ETag       string `xml:"ETag"       json:"etag"`
}

type CompleteMultipartUploadRequest struct {
	Parts []CompletedPart `xml:"Part" json:"parts"`
}
