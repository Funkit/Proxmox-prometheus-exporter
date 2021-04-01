package main

import (
	"fmt"
	"proxmox-prometheus-exporter/connectioninfo"
)

func main() {
	var connInfo connectioninfo.ConnectionInfo

	connInfo.ReadFile("secrets.yml")

	fmt.Println(connInfo)
}
