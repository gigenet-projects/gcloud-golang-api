package api

import (
	"encoding/xml"
	"fmt"
)

type Error struct {
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}

type errorResponse struct {
	Errors []*Error `xml:"Errors>Error"`
}

func ProcessError(b []byte) (err error) {
	body := &errorResponse{}
	err = xml.Unmarshal(b, &body)
	if err != nil || len(body.Errors) == 0 {
		return nil
	}

	return fmt.Errorf("server error [%s] %s", body.Errors[0].Code, body.Errors[0].Message)
}
