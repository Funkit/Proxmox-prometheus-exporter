package exporter

import (
	"fmt"
	"log"
	"net/http"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuLoad = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "PVE",
			Subsystem: "hardware",
			Name:      "cpu_current",
			Help:      "Current CPU load.",
		},
		[]string{
			// name of the hypervisor, VM or container
			"name",
			// Of what type is the element: host, VM or container
			"type",
			// matching resource pool
			"pool",
		},
	)
)

func init() {
	prometheus.MustRegister(cpuLoad)
}

func recordMetrics(accessInfo *connection.Info, conf *Configuration) {
	go func() {
		client := api.NewClient(accessInfo)
		for {
			//output, err := client.GetRawResponse("/cluster/resources")
			//if err != nil {
			//	log.Fatal(err)
			//}
			//fmt.Println("RAW:", output)
			nodeList, vmList, err := client.GetClusterResources()
			if err != nil {
				log.Fatal(err)
			}
			for _, vm := range vmList {
				fmt.Println("target", vm.Name, "has CPU", vm.CPU)
				cpuLoad.WithLabelValues(vm.Name, "VM", vm.Pool).Add(vm.CPU)
			}
			for _, node := range nodeList {
				fmt.Println("target", node.Node, "has CPU", node.CPU)
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

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+configuration.ExposedPort, nil)
}
