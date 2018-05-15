package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) AssociateAddress(publicIp, instanceId string) (err error) {
	if publicIp == "" {
		return errors.New("publicIp is required")
	}
	if instanceId == "" {
		return errors.New("instanceId is required")
	}

	params := fmt.Sprintf("PublicIp=%s&InstanceId=%s", publicIp, instanceId)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "AssociateAddress", params)
	if err != nil {
		return
	}

	err = ProcessError(b)
	return
}
