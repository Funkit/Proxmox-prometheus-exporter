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

	// querying /cluster/resources
	respBody, err := client.Get("/cluster/resources")
	if err != nil {
		log.Fatal(err)
	}

	var buffer Results
	if err := json.Unmarshal([]byte(respBody), &buffer); err != nil {
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

		respBody, err := client.Get("/nodes/" + node.Node + "/network")
		if err != nil {
			log.Fatal(err)
		}

		var buffer map[string]interface{}

		if err := json.Unmarshal([]byte(respBody), &buffer); err != nil {
			log.Fatal(err)
		}

		fmt.Println(buffer)
	}
}
