package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func BuildURL(mode *string, host *string, maxNodes *int, nodes *[]string) (*url.URL, error) {
	url, err := url.Parse(fmt.Sprintf("https://check-host.net/check-%s", *mode))
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Add("host", *host)
	query.Add("max_nodes", strconv.Itoa(*maxNodes))

	if len(*nodes) >= 1 {
		for _, n := range *nodes {
			trim := strings.TrimSpace(n)
			if trim == "" {
				continue
			} else {
				query.Add("node", trim)
			}
		}
	}

	url.RawQuery = query.Encode()
	return url, nil
}
