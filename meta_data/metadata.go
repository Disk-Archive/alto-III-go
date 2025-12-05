package meta_data

import "encoding/xml"

type (
	MetaData struct {
		Object *Object
	}
	S3Tag struct {
		Key   string `xml:"Key" json:"key"`
		Value string `xml:"Value" json:"value"`
	}

	S3TagSet struct {
		Tags []S3Tag `xml:"Tag" json:"tags"`
	}

	S3Tagging struct {
		XMLName xml.Name `xml:"Tagging" json:"tagging"`
		TagSet  S3TagSet `xml:"TagSet" json:"tag_set"`
	}
)
