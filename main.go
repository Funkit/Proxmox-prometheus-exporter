package main

import "github.com/Funkit/proxmox-prometheus-exporter/exporter"

const secretsFilePath = "../secrets/secrets_perso.yml"
const configurationFilePath = "configuration.yml"

func main() {
	exporter.RegisterMetrics()
	exporter.ServeMetrics(secretsFilePath, configurationFilePath)
}
