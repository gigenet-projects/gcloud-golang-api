package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

func (api *Client) CreateKeyPair(keyname string) (key *KeyPair, err error) {
	if keyname == "" {
		return nil, errors.New("keyname is required")
	}

	params := fmt.Sprintf("KeyName=%s", keyname)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "CreateKeyPair", params)
	if err != nil {
		return
	}

	body := &KeyPair{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	key = body
	err = ProcessError(b)

	return
}
