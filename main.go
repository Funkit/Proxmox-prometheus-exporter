package main

import (
	"fmt"
	"log"
	"net/http"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const SECRETS_FILE_PATH = "../secrets/secrets_perso.yml"

func main() {
	var connInfo connection.Info
	connInfo.ReadFile(SECRETS_FILE_PATH)

	client := api.NewClient(&connInfo)

	// /nodes
	nodes, err := client.GetNodes()
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		fmt.Println("Node name:" + node.Name + "; Fingerprint:" + node.SslFingerprint)
	}

	// /cluster/resources

	nodeList, vmList, err := client.GetClusterResources()
	if err != nil {
		log.Fatal(err)
	}

	for _, vm := range vmList {
		fmt.Println("VMID:" + strconv.Itoa(vm.VMID) + "; Status:" + vm.Status)
	}

	for _, node := range nodeList {
		fmt.Println("Node:" + node.Node + "; Status:" + node.Status)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
