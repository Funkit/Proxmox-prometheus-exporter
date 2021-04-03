package main

import (
	"encoding/json"
	"fmt"
	"log"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"strconv"
)

const SECRETS_FILE_PATH = "../secrets/secrets_perso.yml"

func main() {
	var connInfo connection.Info
	connInfo.ReadFile(SECRETS_FILE_PATH)

	client := api.NewClient(&connInfo)

	// /nodes
	respBody1, err := client.Get("/nodes")
	if err != nil {
		log.Fatal(err)
	}

	var buffer1 api.Results
	if err := json.Unmarshal([]byte(respBody1), &buffer1); err != nil {
		log.Fatal(err)
	}

	nodes, err := api.ParseNodes(respBody1)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		fmt.Println("Node name:" + node.Name + "; Fingerprint:" + node.SslFingerprint)
	}

	// /cluster/resources
	respBody, err := client.Get("/cluster/resources")
	if err != nil {
		log.Fatal(err)
	}

	var buffer2 api.Results
	if err := json.Unmarshal([]byte(respBody), &buffer2); err != nil {
		log.Fatal(err)
	}

	nodeList, vmList, err := api.ParseClusterResources(respBody)
	if err != nil {
		log.Fatal(err)
	}

	for _, vm := range vmList {
		fmt.Println("VMID:" + strconv.Itoa(vm.VMID) + "; Status:" + vm.Status)
	}

	for _, node := range nodeList {
		fmt.Println("Node:" + node.Node + "; Status:" + node.Status)
	}
}
