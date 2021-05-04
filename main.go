package main

import "github.com/Funkit/proxmox-prometheus-exporter/exporter"

const configurationFilePath = "./configuration.yml"

func main() {
	exporter.RegisterMetrics()
	exporter.ServeMetrics(configurationFilePath)
}
