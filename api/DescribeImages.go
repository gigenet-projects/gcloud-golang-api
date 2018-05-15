package api

import (
	"encoding/xml"
	"time"
)

type Image struct {
	Id             string `xml:"imageId"`
	Location       string `xml:"imageLocation"`
	State          string `xml:"imageState"`
	Owner          string `xml:"imageOwnerId"`
	Public         bool   `xml:"isPublic"`
	Architecture   string `xml:"architecture"`
	Type           string `xml:"imageType"`
	OwnerAlias     string `xml:"imageOwnerAlias"`
	Name           string `xml:"name"`
	Description    string `xml:"description"`
	Platform       string `xml:"platform"`
	Virtualization string `xml:"virtualizationType"`
	Hypervisor     string `xml:"hypervisor"`
}

type describeImagesResponse struct {
	Images []*Image `xml:"imagesSet>item"`
}

func (api *Client) DescribeImages(filter string) (images []*Image, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeImages", filter)
	if err != nil {
		return
	}

	body := &describeImagesResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	images = body.Images
	return
}
