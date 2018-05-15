package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) RebootInstances(instanceId ...string) (err error) {
	if len(instanceId) == 0 {
		return errors.New("instanceId is required")
	}

	var params string
	for i, id := range instanceId {
		params = fmt.Sprintf("%s&InstanceId.%d=%s", params, i+1, id)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(),
		"RebootInstances", params[1:])

	if err == nil {
		err = ProcessError(b)
	}

	return
}
