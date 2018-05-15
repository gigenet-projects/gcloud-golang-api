package api

import (
	"encoding/xml"
	"fmt"
	"time"
)

func (api *Client) DescribeVmTypes(instanceType string) (types []*InstanceType, err error) {
	var params string
	if instanceType != "" {
		params = fmt.Sprintf("InstanceType=%s", instanceType)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"DescribeVmTypes", params)
	if err != nil {
		return
	}

	body := &describeInstanceTypesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	types = body.Types
	return
}
