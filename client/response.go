package client

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Code     string `json:"code"`
	Country  string `json:"country"`
	City     string `json:"city"`
	IP       string `json:"ip"`
	ASNumber string `json:"as_number"`
}

// ["us", "USA", "Los Angeles", "5.253.30.82", "AS18978"]
func (n *Node) UnmarshalJSON(data []byte) error {
	var raw []string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 5 {
		return fmt.Errorf("expected 5 elements, got %d", len(raw))
	}

	n.Code = raw[0]
	n.Country = raw[1]
	n.City = raw[2]
	n.IP = raw[3]
	n.ASNumber = raw[4]
	return nil
}

type CheckResponse struct {
	Ok            int             `json:"ok"`
	RequestID     string          `json:"request_id"`
	PermanentLink string          `json:"permanent_link"`
	Nodes         map[string]Node `json:"nodes"`
}
