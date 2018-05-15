package api

import (
	"sort"
	"strings"
	"time"
)

func buildQueryString(action string, date time.Time, params ...string) string {
	if action == "" {
		return ""
	}

	p := []string{
		"Action=" + action,
		"Version=2009-03-31",
	}

	// append parameters
	for _, l := range params {
		p = append(p, l)
	}

	// signature expects byte-ordered keys
	sort.Strings(p)

	// return as a series of urlencoded parameters
	return strings.Join(p, "&")
}
