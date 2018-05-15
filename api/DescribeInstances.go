package api

import (
	"encoding/xml"
	"time"
)

type Instance struct {
	// Not implemented:
	// - Placement.AvailabilityZone

	ImageId          string        `xml:"imageId"`
	Id               string        `xml:"instanceId"`
	Type             string        `xml:"instanceType"`
	LaunchTime       time.Time     `xml:"launchTime"`
	PrivateDnsName   string        `xml:"privateDnsName"`
	PrivateIpAddress string        `xml:"privateIpAddress"`
	PublicIpAddress  string        `xml:"publicIpAddress"`
	State            InstanceState `xml:"state"`
	SubnetId         string        `xml:"subnetId"`
	Architecture     string        `xml:"architecture"`
	Virtualization   string        `xml:"virtualizationType"`
	RootDeviceName   string        `xml:"rootDeviceName"`
	Hypervisor       string        `xml:"hypervisor"`
	Platform         string        `xml:"platform"`
	Tags             []*Tag        `xml:"tagSet>item"`
}

type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

type Reservation struct {
	OwnerId   int
	Instances []*Instance `xml:"instancesSet>item"`
}

type describeInstancesResponse struct {
	Reservations []*Reservation `xml:"reservationSet>item"`
}

func (api *Client) DescribeInstances(filter string) (reservations []*Reservation, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeInstances", filter)
	if err != nil {
		return
	}

	body := &describeInstancesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	reservations = body.Reservations
	return
}
