package api

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

func (api *Client) ImportKeyPair(keyname string, public string) (key *KeyPair, err error) {
	if keyname == "" {
		return nil, errors.New("keyname is required")
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(public))
	params := fmt.Sprintf("KeyName=%s&PublicKeyMaterial=%s", keyname, b64)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "ImportKeyPair", params)
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
