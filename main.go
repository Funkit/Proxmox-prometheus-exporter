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

type Results struct {
	Rows []map[string]interface{} `json:"data"`
}

func main() {
	var connInfo connection.Info
	connInfo.ReadFile(SECRETS_FILE_PATH)

	client := api.NewClient(&connInfo)

	// querying /nodes
	respBody, err := client.Get("/nodes")
	if err != nil {
		log.Fatal(err)
	}

	var queryOutput api.Nodes
	if err := json.Unmarshal([]byte(respBody), &queryOutput); err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryOutput)

	// querying /cluster/resources
	respBody2, err := client.Get("/cluster/resources")
	if err != nil {
		log.Fatal(err)
	}

	var buffer Results
	if err := json.Unmarshal([]byte(respBody2), &buffer); err != nil {
		log.Fatal(err)
	}

	nodeList, vmList, err := api.ParseClusterResources(respBody2)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodeList {
		fmt.Println("Node:" + node.Node + "; Status:" + node.Status)
	}
	for _, vm := range vmList {
		fmt.Println("VMID:" + strconv.Itoa(vm.VMID) + "; Status:" + vm.Status)
	}
}
