package main

import "proxmox-prometheus-exporter/exporter"

const secretsFilePath = "../secrets/secrets_perso.yml"
const configurationFilePath = "configuration.yml"

func main() {
	exporter.ServeMetrics(secretsFilePath, configurationFilePath)
}
