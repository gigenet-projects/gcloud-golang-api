package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) UnassignPublicIpAddress(publicIp string) (err error) {
	if publicIp == "" {
		return errors.New("publicIp is required")
	}

	params := fmt.Sprintf("PublicIp=%s", publicIp)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "UnassignPublicIpAddress", params)
	if err != nil {
		return
	}

	err = ProcessError(b)
	return
}
