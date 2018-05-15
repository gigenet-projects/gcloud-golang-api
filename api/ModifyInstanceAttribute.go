package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) ModifyInstanceAttribute(attribute, instanceId string, extras []string) (err error) {
	if instanceId == "" {
		return errors.New("instanceId is required")
	}

	var params string
	if attribute != "" {
		params = fmt.Sprintf("Attribute=%s", attribute)
	}
	if params != "" {
		params += "&"
	}
	params = fmt.Sprintf("%sInstanceId=%s", params, instanceId)
	for _, v := range extras {
		params = fmt.Sprintf("%s&%s", params, v)
	}

	_, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"ModifyInstanceAttribute", params)
	return
}
