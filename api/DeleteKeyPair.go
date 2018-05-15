package api

import (
	"errors"
	"fmt"
	"time"
)

func (api *Client) DeleteKeyPair(keyname string) (err error) {
	if keyname == "" {
		return errors.New("keyname is required")
	}

	params := fmt.Sprintf("KeyName=%s", keyname)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "DeleteKeyPair", params)
	if err != nil {
		return
	}

	err = ProcessError(b)
	return
}
