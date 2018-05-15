package api

import (
	"encoding/xml"
	"time"
)

type Address struct {
	PublicIp   string `xml:"publicIp",omitempty`
	PrivateIp  string `xml:"privateIpAddress",omitempty`
	Domain     string `xml:"domain"`
	InstanceId string `xml:"instanceId"`
}

type describeAddressesResponse struct {
	Addresses []*Address `xml:"addressesSet>item"`
}

func (api *Client) DescribeAddresses(filter string) (addresses []*Address, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeAddresses", filter)
	if err != nil {
		return
	}

	body := &describeAddressesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	addresses = body.Addresses
	return
}
