package main

import (
	"fmt"
	"proxmox-prometheus-exporter/connectioninfo"
)

const SECRETS_FILE_PATH = "../secrets/secrets.yml"

func main() {
	var connInfo connectioninfo.ConnectionInfo

	connInfo.ReadFile(SECRETS_FILE_PATH)

	fmt.Println(connInfo)
}
