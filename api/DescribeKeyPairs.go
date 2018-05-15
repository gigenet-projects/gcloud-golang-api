package api

import (
	"encoding/xml"
	"time"
)

type KeyPair struct {
	Name        string `xml:"keyName"`
	Fingerprint string `xml:"keyFingerprint"`
	Material    string `xml:"keyMaterial",omitempty`
}

type describeKeyPairsResponse struct {
	Keys []*KeyPair `xml:"keySet>item"`
}

func (api *Client) DescribeKeyPairs(filter string) (keys []*KeyPair, err error) {
	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DescribeKeyPairs", filter)
	if err != nil {
		return
	}

	body := &describeKeyPairsResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	keys = body.Keys
	return
}
