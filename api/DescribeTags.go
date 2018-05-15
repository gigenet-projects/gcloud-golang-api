package api

import (
	"encoding/xml"
	"time"
)

type Tag struct {
	ResourceId   string `xml:"resourceId"`
	ResourceType string `xml:"resourceType"`
	Key          string `xml:"key"`
	Value        string `xml:"omitempty"`
}

type describeTagsResponse struct {
	Tags []*Tag `xml:"tagSet>item"`
}

func (api *Client) DescribeTags(filter string) (tags []*Tag, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeTags", filter)
	if err != nil {
		return
	}

	body := &describeTagsResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	tags = body.Tags
	return
}
