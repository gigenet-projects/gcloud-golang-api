package api

import (
	"encoding/xml"
	"fmt"
	"time"
)

type SessionToken struct {
	SessionToken    string
	SecretAccessKey string
	Expiration      time.Time
	AccessKeyId     string
}

type getSessionTokenResponse struct {
	GetSessionTokenResult struct {
		Credentials []SessionToken `xml:"Credentials>item"`
	}
}

func (api *Client) GetSessionToken(durationSeconds int) (token *SessionToken, err error) {
	if durationSeconds < 900 ||
		durationSeconds > 129600 {
		return token, fmt.Errorf("durationSeconds must be > 900 and < 129600")
	}

	var b []byte
	p := fmt.Sprintf("DurationSeconds=%d", durationSeconds)
	b, err = Request(api.server, api.ApiKey, api.secret, time.Now(), "GetSessionToken", p)
	if err != nil {
		return
	}

	body := &getSessionTokenResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil {
		return
	}

	if len(body.GetSessionTokenResult.Credentials) == 0 {
		return token, fmt.Errorf("server did not return any credentials")
	}

	token = &body.GetSessionTokenResult.Credentials[0]
	return
}
