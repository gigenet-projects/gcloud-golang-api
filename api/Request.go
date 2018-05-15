package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	awsauth "github.com/smartystreets/go-aws-auth"
)

func Request(BASEURL, APIKEY, SECRET string, date time.Time, action, params string) (data []byte, err error) {
	if BASEURL == "" {
		return nil, errors.New("baseurl cannot be empty")
	}
	if APIKEY == "" {
		return nil, errors.New("apikey cannot be empty")
	}
	if SECRET == "" {
		return nil, errors.New("secret cannot be empty")
	}
	if action == "" {
		return nil, errors.New("action cannot be empty")
	}

	u, err := url.Parse(BASEURL)
	if err != nil {
		return nil, fmt.Errorf("invalid BASEURL: %s", err)
	}

	if u.Path == "" {
		u.Path = "/"
	}

	qs := buildQueryString(action, date, params)
	url := fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.Path, qs)

	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not form request: %s", err)
	}
	awsauth.Sign2(req, awsauth.Credentials{
		AccessKeyID:     APIKEY,
		SecretAccessKey: SECRET,
	})

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}
