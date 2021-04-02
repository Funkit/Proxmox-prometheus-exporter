package main

import (
	"encoding/json"
	"fmt"
	"log"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
)

const SECRETS_FILE_PATH = "../secrets/secrets.yml"

func main() {
	var connInfo connection.Info
	connInfo.ReadFile(SECRETS_FILE_PATH)

	client := api.NewClient(&connInfo)

	respBody, err := client.Get("/nodes")
	if err != nil {
		log.Fatal(err)
	}

	var queryOutput api.Nodes
	if err := json.Unmarshal([]byte(respBody), &queryOutput); err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryOutput)
}
