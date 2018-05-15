package api

import (
	"encoding/xml"
	"time"
)

type ReservedInstance struct {
	Id                 string    `xml:"reservedInstancesId"`
	InstanceType       string    `xml:"instanceType"`
	AvailabilityZone   string    `xml:"availabilityZone"`
	UsagePrice         float64   `xml:"usagePrice"`
	InstanceCount      int       `xml:"instanceCount"`
	ProductDescription string    `xml:"productDescription"`
	Start              time.Time `xml:"start"`
	State              string    `xml:"state"`
	InstanceTenancy    string    `xml:"instanceTenancy"`
	CurrencyCode       string    `xml:"currencyCode"`
	OfferingType       string    `xml:"offeringType"`
}

type describeReservedInstancesResponse struct {
	Reservations []*ReservedInstance `xml:"reservedInstancesSet>item"`
}

func (api *Client) DescribeReservedInstances(filter string) (reservations []*ReservedInstance, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeReservedInstances", filter)
	if err != nil {
		return
	}

	body := &describeReservedInstancesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	reservations = body.Reservations
	return
}
