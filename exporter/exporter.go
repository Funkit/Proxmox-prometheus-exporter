package exporter

import (
	"log"
	"net/http"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.MustRegister(cpuLoad)
}

func recordMetrics(accessInfo *connection.Info, conf *Configuration) {
	go func() {
		client := api.NewClient(accessInfo)
		for {
			nodeList, vmList, err := client.GetClusterResources()
			if err != nil {
				log.Fatal(err)
			}
			for _, vm := range vmList {
				cpuLoad.WithLabelValues(vm.Name, "VM", vm.Pool).Add(vm.CPU)
				maxCpu.WithLabelValues(vm.Name, "VM", vm.Pool).Add(float64(vm.AllocatedCPUCores))
				ramLoad.WithLabelValues(vm.Name, "VM", vm.Pool).Add(float64(vm.RAM))
				maxRam.WithLabelValues(vm.Name, "VM", vm.Pool).Add(float64(vm.AllocatedRAMBytes))
			}
			for _, node := range nodeList {
				cpuLoad.WithLabelValues(node.Node, "Node", "N/A").Add(node.CPU)
			}

			time.Sleep(time.Duration(conf.QueryPeriod) * time.Second)
		}
	}()
}

//ServeMetrics main HTTP server for Prometheus metrics
func ServeMetrics(secretsFilePath string, configurationFilePath string) {
	configuration, err := GetConfigurationFromFile(configurationFilePath)
	if err != nil {
		log.Fatal(err)
	}

	connectionInfo, err := connection.GetInfoFromFile(secretsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	recordMetrics(connectionInfo, configuration)

	http.Handle(configuration.MetricsPath, promhttp.Handler())
	http.ListenAndServe(":"+configuration.ExposedPort, nil)
}
