package exporter

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/Funkit/pve-go-api/api"
	apiconnection "github.com/Funkit/pve-go-api/connection"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//RegisterMetrics register all Prometheus metrics
func RegisterMetrics() {
	prometheus.MustRegister(cpuLoad)
	prometheus.MustRegister(maxCPU)
	prometheus.MustRegister(ramLoad)
	prometheus.MustRegister(maxRAM)
	prometheus.MustRegister(uptimeSec)
}

func exportClusterResources(client *api.Client) {
	resList, err := client.GetClusterResources()
	if err != nil {
		panic(err)
	}
	for _, res := range resList {

		if res.Type == "node" {
			cpuLoad.WithLabelValues(res.Node, res.Type, "N/A").Add(res.CPU)
			maxCPU.WithLabelValues(res.Node, res.Type, "N/A").Add(float64(res.AllocatedCPU))
			ramLoad.WithLabelValues(res.Node, res.Type, "N/A").Add(float64(res.RAM))
			maxRAM.WithLabelValues(res.Node, res.Type, "N/A").Add(float64(res.AllocatedRAMBytes))
			uptimeSec.WithLabelValues(res.Node, res.Type, "N/A").Add(float64(res.Uptime))
		}
		if res.Type == "qemu" {
			cpuLoad.WithLabelValues(res.Name, res.Type, emptyIsNA(res.Pool)).Add(res.CPU)
			maxCPU.WithLabelValues(res.Name, res.Type, emptyIsNA(res.Pool)).Add(float64(res.AllocatedCPU))
			ramLoad.WithLabelValues(res.Name, res.Type, emptyIsNA(res.Pool)).Add(float64(res.RAM))
			maxRAM.WithLabelValues(res.Name, res.Type, emptyIsNA(res.Pool)).Add(float64(res.AllocatedRAMBytes))
			uptimeSec.WithLabelValues(res.Name, res.Type, emptyIsNA(res.Pool)).Add(float64(res.Uptime))
		}

	}
}

func emptyIsNA(entry string) string {
	if entry == "" {
		return "N/A"
	}
	return entry
}

func recordMetrics(conf Configuration) {
	go func() {
		connectionInfo, err := apiconnection.ReadFile(conf.SecretsPath)
		if err != nil {
			panic(err)
		}

		fmt.Println("Connecting to PVE at", connectionInfo.Address)

		client := api.NewClient(*connectionInfo, &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: true,
		})
		for {
			exportClusterResources(client)
			time.Sleep(time.Duration(conf.QueryPeriod) * time.Second)
		}
	}()
}

//ServeMetrics main HTTP server for Prometheus metrics
func ServeMetrics(configurationFilePath string) {
	configuration, err := GetConfigurationFromFile(configurationFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Metrics path:", configuration.MetricsPath, "; exposed port:", configuration.ExposedPort)

	recordMetrics(configuration)

	http.Handle(configuration.MetricsPath, promhttp.Handler())
	http.ListenAndServe(":"+configuration.ExposedPort, nil)
}
