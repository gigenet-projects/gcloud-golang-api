package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

func (api *Client) AssignPublicIpAddress(instanceId string) (address *Address, err error) {
	if instanceId == "" {
		return nil, errors.New("instanceId is required")
	}

	params := fmt.Sprintf("InstanceId=%s", instanceId)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "AssignPublicIpAddress", params)
	if err != nil {
		return
	}

	body := &Address{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	address = body
	return
}
