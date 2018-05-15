package api

import (
	"fmt"
	"time"
)

func (api *Client) DescribeImageAttribute(attribute, imageId string) (b []byte, err error) {
	var params string
	if attribute != "" {
		params = fmt.Sprintf("Attribute=%s", attribute)
	}
	if imageId != "" {
		if params != "" {
			params += "&"
		}
		params = fmt.Sprintf("%sImageId=%s", params, imageId)
	}

	return Request(api.server, api.ApiKey, api.secret, time.Now(),
		"DescribeImageAttribute", params)
}
