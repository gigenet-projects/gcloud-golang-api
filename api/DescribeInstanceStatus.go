package api

import (
	"fmt"
	"time"
)

func (api *Client) DescribeInstanceStatus(instanceId string) (b []byte, err error) {
	var params string
	if instanceId != "" {
		params = fmt.Sprintf("InstanceId=%s", instanceId)
	}

	return Request(api.server, api.ApiKey, api.secret, time.Now(),
		"DescribeInstanceStatus", params)
}
