package api

import (
	"fmt"
	"time"
)

func (api *Client) DescribeInstanceAttribute(attribute, instanceId string) (b []byte, err error) {
	var params string
	if attribute != "" {
		params = fmt.Sprintf("Attribute=%s", attribute)
	}
	if instanceId != "" {
		if params != "" {
			params += "&"
		}
		params = fmt.Sprintf("%sInstanceId=%s", params, instanceId)
	}

	return Request(api.server, api.ApiKey, api.secret, time.Now(),
		"DescribeInstanceAttribute", params)
}
