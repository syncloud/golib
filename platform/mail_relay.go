package platform

import (
	"encoding/json"
	"fmt"
	"io"
)

type MailRelay struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type MailRelayResponse struct {
	Success bool      `json:"success"`
	Data    MailRelay `json:"data"`
}

func (c *Client) GetMailRelay() (*MailRelay, error) {
	c.logger.Info("get mail relay")
	resp, err := c.client.Get("http://unix/mail/relay")
	if err != nil {
		return nil, err
	}
	// a platform that predates the mail relay has no such endpoint, which means
	// the feature is simply unavailable rather than broken
	if resp.StatusCode == 404 {
		return &MailRelay{Enabled: false}, nil
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get mail relay, %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var responseJson MailRelayResponse
	if err := json.Unmarshal(bodyBytes, &responseJson); err != nil {
		return nil, err
	}
	return &responseJson.Data, nil
}
