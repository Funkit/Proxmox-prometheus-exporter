package main

import (
	"fmt"
	"log"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"strconv"
)

const secretsFilePath = "../secrets/secrets_perso.yml"

func main() {
	connInfo, err := connection.GetInfoFromFile(secretsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewClient(connInfo)

	// /nodes
	nodes, err2 := client.GetNodes()
	if err != nil {
		log.Fatal(err2)
	}

	for _, node := range nodes {
		fmt.Println("Node name:" + node.Name + "; Fingerprint:" + node.SslFingerprint)
	}

	// /cluster/resources

	nodeList, vmList, err3 := client.GetClusterResources()
	if err3 != nil {
		log.Fatal(err3)
	}

	for _, vm := range vmList {
		fmt.Println("VMID:" + strconv.Itoa(vm.VMID) + "; Status:" + vm.Status)
	}

	for _, node := range nodeList {
		fmt.Println("Node:" + node.Node + "; Status:" + node.Status)

		networkInterfaceList, err4 := client.GetNodeNetwork(node.Node)
		if err4 != nil {
			log.Fatal(err4)
		}
		for _, networkInterface := range networkInterfaceList {
			fmt.Println(networkInterface)
		}
	}
}
