package api

import (
	"encoding/xml"
	"time"
)

type AvailabilityZone struct {
	Region   string   `xml:"regionName"`
	Name     string   `xml:"zoneName"`
	State    string   `xml:"zoneState"`
	Messages []string `xml:"messageSet>message"`
}

type describeAvailabilityZonesResponse struct {
	AvailabilityZones []*AvailabilityZone `xml:"availabilityZoneInfo>item"`
}

func (api *Client) DescribeAvailabilityZones(filter string) (availabilityZones []*AvailabilityZone, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeAvailabilityZones", filter)
	if err != nil {
		return
	}

	body := &describeAvailabilityZonesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	availabilityZones = body.AvailabilityZones
	return
}
