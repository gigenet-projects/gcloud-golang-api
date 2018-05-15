package api

import (
	"fmt"
	"time"
)

func (api *Client) CreateTags(tags []*Tag, instanceId string) (err error) {
	params := fmt.Sprintf("ResourceId.1=%s", instanceId)
	for i, t := range tags {
		params = fmt.Sprintf("%s&Tag.%d.Key=%s&Tag.%d.Value=%s", params, i+1, t.Key, i+1, t.Value)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "CreateTags", params)
	if err != nil {
		return
	}

	err = ProcessError(b)
	return
}
