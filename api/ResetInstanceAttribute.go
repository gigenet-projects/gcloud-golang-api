package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) ResetInstanceAttribute(attribute, instanceId string) (err error) {
	if attribute == "" {
		return errors.New("attribute is required")
	}
	if instanceId == "" {
		return errors.New("instanceId is required")
	}

	params := fmt.Sprintf("Attribute=%s&InstanceId=%s", attribute, instanceId)
	_, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"ResetInstanceAttribute", params)
	return
}
