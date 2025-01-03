package utils

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func ValidateMode(mode *string) error {
	if len(*mode) <= 0 {
		return errors.New("string must not be empty")
	}

	s := strings.ToLower(*mode)
	if s != "http" && s != "tcp" && s != "udp" && s != "dns" {
		return fmt.Errorf("want 'http', 'tcp', 'udp' or 'dns', got '%s' instead", s)
	}

	return nil
}

func ValidateMaxNodes(m *int) error {
	if *m <= 0 {
		return errors.New("integer must be greater than zero")
	}

	return nil
}

func ParseNodeList(l *string, maxNodes *int) ([]string, error) {
	n := strings.Split(*l, ",")
	c := []string{}

	// make sure nodes aren't just an empty space
	for _, i := range n {
		trim := strings.TrimSpace(i)
		if trim == "" {
			continue
		} else {
			c = append(c, i)
		}
	}

	if len(c) > *maxNodes {
		return nil, errors.New("the amount of provided nodes exceeds the maximum permitted amount. increase it by using '--max-nodes' or remove the additional nodes")
	}

	return n, nil
}

func GetTarget() (string, error) {
	h := flag.Arg(0)
	if len(h) <= 0 {
		return "", errors.New("a valid, non-empty string must be provided")
	}

	return h, nil
}
