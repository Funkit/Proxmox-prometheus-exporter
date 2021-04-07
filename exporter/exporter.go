package exporter

import (
	"fmt"
	"log"
	"proxmox-prometheus-exporter/api"
	"proxmox-prometheus-exporter/connection"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	cpuLoad = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "PVE",
		Subsystem: "hardware",
		Name:      "cpu_current",
		Help:      "Current CPU load.",
	})
)

func init() {
	prometheus.MustRegister(cpuLoad)
}

func recordMetrics(accessInfo *connection.Info, conf *Configuration) {
	//go func() {
	client := api.NewClient(accessInfo)
	counter := 0
	for {
		nodeList, vmList, err := client.GetClusterResources()
		if err != nil {
			log.Fatal(err)
		}
		for _, vm := range vmList {
			fmt.Printf("%v:%v\n", counter, vm)
		}
		for _, node := range nodeList {
			fmt.Printf("%v:%v\n", counter, node)
		}
		counter++
		opsProcessed.Inc()
		time.Sleep(time.Duration(conf.QueryPeriod) * time.Second)
	}
	//}()
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

	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(configuration.ExposedPort, nil)
}
