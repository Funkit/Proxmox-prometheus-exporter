package api

import (
	"crypto/tls"
	"net/http"
	"proxmox-prometheus-exporter/connection"
)

func TokenHeader(c *connection.Info) string {
	return "PVEAPIToken=" + c.UserId.Username + "@" + c.UserId.IdRealm + "!" + c.ApiToken.Id + "=" + c.ApiToken.Token
}

func NewRequest(c *connection.Info, targetUrl string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.Address+targetUrl, nil)
	req.Header.Add("Authorization", TokenHeader(c))
	return req, err
}

func NewClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client
}
