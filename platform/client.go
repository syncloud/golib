package platform

import (
	"encoding/json"
	"fmt"
	"github.com/syncloud/golib/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
)

type HttpClient interface {
	Post(url string, values url.Values) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
}

type Client struct {
	client HttpClient
	logger *zap.Logger
}

func New() *Client {
	return &Client{
		client: NewHttpClient(),
		logger: log.Logger(),
	}
}

func (c *Client) InitStorage(app, user string) (string, error) {
	values := url.Values{"app_name": {app}, "user_name": {user}}
	c.logger.Info("init storage", zap.String("request", values.Encode()))
	resp, err := c.client.Post("http://unix/app/init_storage", values)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("unable to init storage, %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var responseJson Response
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		return "", err
	}
	return responseJson.Data, nil
}

func (c *Client) GetAppDomainName(app string) (string, error) {
	c.logger.Info("get app domain name", zap.String("app", app))
	resp, err := c.client.Get(fmt.Sprintf("http://unix/app/domain_name?name=%s", app))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("get app domain name, %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var responseJson Response
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		return "", err
	}
	return responseJson.Data, nil
}

func (c *Client) GetAppUrl(app string) (string, error) {
	c.logger.Info("get app url", zap.String("app", app))
	resp, err := c.client.Get(fmt.Sprintf("http://unix/app/url?name=%s", app))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("get app url, %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var responseJson Response
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		return "", err
	}
	return responseJson.Data, nil
}

func (c *Client) RegisterOIDCClient(id, redirectUri string, requirePkce bool, tokenEndpointAuthMethod string) (string, error) {
	requirePkceString := "false"
	if requirePkce {
		requirePkceString = "true"
	}
	values := url.Values{
		"id":                         {id},
		"redirect_uri":               {redirectUri},
		"require_pkce":               {requirePkceString},
		"token_endpoint_auth_method": {tokenEndpointAuthMethod},
	}
	c.logger.Info("register oidc client", zap.String("request", values.Encode()))
	resp, err := c.client.Post("http://unix/oidc/register", values)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("register oidc client, %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var responseJson Response
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		return "", err
	}
	return responseJson.Data, nil
}

func (c *Client) RestartService(name string) error {
	values := url.Values{"name": {name}}
	c.logger.Info("service restart", zap.String("request", values.Encode()))
	resp, err := c.client.Post("http://unix/service/restart", values)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unable to restart service, %s", resp.Status)
	}
	return nil
}
