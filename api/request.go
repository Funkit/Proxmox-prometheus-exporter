package api

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"proxmox-prometheus-exporter/connection"
)

type Client struct {
	httpClient *http.Client
	info       *connection.Info
}

func tokenHeader(c *connection.Info) string {
	return "PVEAPIToken=" + c.UserId.Username + "@" + c.UserId.IdRealm + "!" + c.ApiToken.Id + "=" + c.ApiToken.Token
}

func newRequest(c *connection.Info, targetUrl string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.Address+targetUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tokenHeader(c))
	return req, nil
}

func NewClient(c *connection.Info) *Client {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
	}
	httpClient := &http.Client{Transport: tr}

	client := &Client{
		httpClient: httpClient,
		info:       c,
	}

	return client
}

func (c *Client) Get(url string) (responseBody []byte, err error) {
	req, err := newRequest(c.info, url)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) GetNodes() ([]Node, error) {

	respBody, err := c.Get("/nodes")
	if err != nil {
		return nil, err
	}

	nodeList, err := parseNodes(respBody)
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

func (c *Client) GetClusterResources() ([]NodeResource, []VMResource, error) {

	respBody, err := c.Get("/cluster/resources")
	if err != nil {
		return nil, nil, err
	}

	nodeList, vmList, err := parseClusterResources(respBody)
	if err != nil {
		log.Fatal(err)
	}

	return nodeList, vmList, nil
}
