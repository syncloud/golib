package platform

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syncloud/golib/log"
)

type MailRelayClientStub struct {
	url  string
	body string
}

func (h *MailRelayClientStub) Get(url string) (resp *http.Response, err error) {
	h.url = url
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(h.body))),
	}, nil
}

func (h *MailRelayClientStub) Post(_ string, _ url.Values) (resp *http.Response, err error) {
	return nil, nil
}

func TestClient_GetMailRelay_Enabled(t *testing.T) {
	httpClient := &MailRelayClientStub{body: `
{
	"success": true,
	"data": {
		"enabled": true,
		"host": "mail-relay.syncloud.it",
		"port": 587,
		"login": "device.syncloud.it",
		"password": "the-token"
	}
}
`}
	client := &Client{client: httpClient, logger: log.Logger()}

	relay, err := client.GetMailRelay()
	assert.NoError(t, err)
	assert.Equal(t, "http://unix/mail/relay", httpClient.url)
	assert.True(t, relay.Enabled)
	assert.Equal(t, "mail-relay.syncloud.it", relay.Host)
	assert.Equal(t, 587, relay.Port)
	assert.Equal(t, "device.syncloud.it", relay.Login)
	assert.Equal(t, "the-token", relay.Password)
}

func TestClient_GetMailRelay_Disabled(t *testing.T) {
	httpClient := &MailRelayClientStub{body: `{"success": true, "data": {"enabled": false}}`}
	client := &Client{client: httpClient, logger: log.Logger()}

	relay, err := client.GetMailRelay()
	assert.NoError(t, err)
	assert.False(t, relay.Enabled)
	assert.Empty(t, relay.Host)
}
