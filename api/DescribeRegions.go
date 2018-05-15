package api

import (
	"encoding/xml"
	"time"
)

type Region struct {
	Name     string `xml:"regionName"`
	Endpoint string `xml:"regionEndpoint"`
}

type describeRegionsResponse struct {
	Regions []*Region `xml:"regionInfo>item"`
}

func (api *Client) DescribeRegions(filter string) (regions []*Region, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeRegions", filter)
	if err != nil {
		return
	}

	body := &describeRegionsResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	regions = body.Regions
	return
}
