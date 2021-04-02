package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
)

const SECRETS_FILE_PATH = "../secrets/secrets.yml"

func main() {
	var connInfo connection.Info
	connInfo.ReadFile(SECRETS_FILE_PATH)

	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
	}
	client := &http.Client{Transport: tr}

	req, err := api.NewRequest(&connInfo, "/nodes")
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var queryOutput api.Nodes
	if err := json.Unmarshal([]byte(respBody), &queryOutput); err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryOutput)
}
