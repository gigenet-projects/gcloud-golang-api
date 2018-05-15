package api

import (
	"encoding/xml"
	"fmt"
	"time"
)

type getPasswordDataResponse struct {
	InstanceId   string `xml:"instanceId"`
	PasswordData string `xml:"passwordData"`
}

func (api *Client) GetPasswordData(instanceId string) (password string, err error) {
	var params string
	if instanceId != "" {
		params = fmt.Sprintf("InstanceId=%s", instanceId)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"GetPasswordData", params)

	body := &getPasswordDataResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	password = body.PasswordData
	return
}
