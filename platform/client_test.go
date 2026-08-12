package platform

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/syncloud/golib/log"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type HttpClientStub struct {
	values url.Values
	url    string
}

func (h *HttpClientStub) Get(url string) (resp *http.Response, err error) {
	h.url = url
	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(`
{
	"success": true,
	"data": "/data/app"
}
`))),
	}, nil
}

func (h *HttpClientStub) Post(url string, values url.Values) (resp *http.Response, err error) {
	h.values = values
	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(bytes.NewReader([]byte(`
{
	"success": true,
	"data": "/data/app"
}
`))),
	}, nil
}

func TestRealHttpClient_Post(t *testing.T) {
	httpClient := &HttpClientStub{}
	client := &Client{
		client: httpClient,
		logger: log.Logger(),
	}
	storage, err := client.InitStorage("app", "user")
	assert.NoError(t, err)
	assert.Contains(t, httpClient.values.Encode(), "app")
	assert.Contains(t, httpClient.values.Encode(), "user")
	assert.Equal(t, "/data/app", storage)

}

func TestClient_RegisterOIDCClient_SendsMultipleRedirectUris(t *testing.T) {
	httpClient := &HttpClientStub{}
	client := &Client{
		client: httpClient,
		logger: log.Logger(),
	}
	_, err := client.RegisterOIDCClient("app", []string{"/cb", "/mobile"}, true, "client_secret_basic")
	assert.NoError(t, err)
	assert.Equal(t, "app", httpClient.values.Get("id"))
	assert.Equal(t, []string{"/cb", "/mobile"}, httpClient.values["redirect_uri"])
	assert.Equal(t, "true", httpClient.values.Get("require_pkce"))
	assert.Equal(t, "client_secret_basic", httpClient.values.Get("token_endpoint_auth_method"))
}

func TestClient_GetAppStorageDir(t *testing.T) {
	httpClient := &HttpClientStub{}
	client := &Client{
		client: httpClient,
		logger: log.Logger(),
	}
	storage, err := client.GetAppStorageDir("app")
	assert.NoError(t, err)
	assert.Contains(t, httpClient.url, "/app/storage_dir?name=app")
	assert.Equal(t, "/data/app", storage)
}

type StatusClientStub struct {
	status int
	values url.Values
	url    string
}

func (h *StatusClientStub) Get(url string) (*http.Response, error) {
	h.url = url
	return &http.Response{StatusCode: h.status, Body: io.NopCloser(bytes.NewReader(nil))}, nil
}

func (h *StatusClientStub) Post(url string, values url.Values) (*http.Response, error) {
	h.url = url
	h.values = values
	return &http.Response{StatusCode: h.status, Body: io.NopCloser(bytes.NewReader(nil))}, nil
}

func TestClient_RegisterMailInbound(t *testing.T) {
	httpClient := &StatusClientStub{status: 200}
	client := &Client{client: httpClient, logger: log.Logger()}

	err := client.RegisterMailInbound("/var/snap/mail/current/spool/public/tunnel")

	assert.NoError(t, err)
	assert.Equal(t, "http://unix/mail/inbound/register", httpClient.url)
	assert.Equal(t, "/var/snap/mail/current/spool/public/tunnel", httpClient.values.Get("socket"))
}

func TestClient_RegisterMailInboundNotSupported(t *testing.T) {
	client := &Client{client: &StatusClientStub{status: 404}, logger: log.Logger()}

	err := client.RegisterMailInbound("/socket")

	assert.ErrorIs(t, err, ErrNotSupported)
}

func TestClient_RegisterMailInboundFails(t *testing.T) {
	client := &Client{client: &StatusClientStub{status: 500}, logger: log.Logger()}

	err := client.RegisterMailInbound("/socket")

	assert.Error(t, err)
	assert.NotErrorIs(t, err, ErrNotSupported)
}
