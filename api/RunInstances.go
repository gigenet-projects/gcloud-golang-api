package api

import (
	"bytes"
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

type runInstancesResponse struct {
	Owner     string      `xml:"ownerId"`
	Instances []*Instance `xml:"instancesSet>item"`
}

func (api *Client) RunInstances(imageId, instanceType string, extras map[string]string) (instances []*Instance, err error) {
	if imageId == "" {
		return nil, errors.New("imageId is required.")
	}
	if instanceType == "" {
		return nil, errors.New("instanceType is required.")
	}

	var buf bytes.Buffer
	buf.Grow(128)
	buf.WriteString("ImageId=")
	buf.WriteString(imageId)
	buf.WriteString("&InstanceType=")
	buf.WriteString(instanceType)
	if len(extras) > 0 {
		for k, v := range extras {
			switch strings.ToLower(k) {
			case "keyname":
				buf.WriteString("&KeyName=")
				buf.WriteString(v)
			// NetworkInterface
			// TagSpecification
			case "privateipaddress":
				buf.WriteString("&PrivateIpAddress=")
				buf.WriteString(v)
			case "mincount":
				buf.WriteString("&MinCount=")
				buf.WriteString(v)
			case "maxcount":
				buf.WriteString("&MaxCount=")
				buf.WriteString(v)
			case "billingterm":
				buf.WriteString("&BillingTerm=")
				buf.WriteString(v)
			case "placement_availabilityzone":
                                buf.WriteString("&Placement_AvailabilityZone=")
                                buf.WriteString(v)
			}
		}
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "RunInstances", buf.String())
	body := &runInstancesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	instances = body.Instances
	err = ProcessError(b)
	return
}
