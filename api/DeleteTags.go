package api

import (
	"fmt"
	"time"
)

func (api *Client) DeleteTags(tags []*Tag, instanceId string) (err error) {
	params := fmt.Sprintf("ResourceId.1=%s", instanceId)
	for i, t := range tags {
		params = fmt.Sprintf("%s&Tag.%d.Key=%s", params, i+1, t.Key)
	}

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DeleteTags", params)
	if err != nil {
		return
	}

	err = ProcessError(b)
	return
}
