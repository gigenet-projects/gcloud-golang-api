package api

import (
	"encoding/xml"
	"fmt"
	"time"
)

type InstanceType struct {
	Name   string `xml:"name"`
	Cores  int    `xml:"cores"`
	Memory int    `xml:"memory"`
	Disks  int    `xml:"disk"`
}

type describeInstanceTypesResponse struct {
	Types []*InstanceType `xml:"instanceTypeSet>item"`
}

func (api *Client) DescribeInstanceTypes(instanceType string) (types []*InstanceType, err error) {
	var params string
	if instanceType != "" {
		params = fmt.Sprintf("InstanceType=%s", instanceType)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"DescribeInstanceTypes", params)
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
