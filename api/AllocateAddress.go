package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

func (api *Client) AllocateAddress(domain string) (address *Address, err error) {
	if domain == "" {
		return nil, errors.New("domain is required")
	}

	params := fmt.Sprintf("Domain=%s", domain)

	var b []byte
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "AllocateAddress", params)
	if err != nil {
		return
	}

	body := &Address{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	address = body
	return
}
