package exporter

import (
	"crypto/tls"
	"log"
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
}

func exportClusterResources(client *api.Client) {
	resList, err := client.GetClusterResources()
	if err != nil {
		log.Fatal(err)
	}
	for _, res := range resList {
		cpuLoad.WithLabelValues(res.Name, "VM", res.Pool).Add(res.CPU)
		maxCpu.WithLabelValues(res.Name, "VM", res.Pool).Add(float64(res.AllocatedCPU))
		ramLoad.WithLabelValues(res.Name, "VM", res.Pool).Add(float64(res.RAM))
		maxRam.WithLabelValues(res.Name, "VM", res.Pool).Add(float64(res.AllocatedRAMBytes))
	}
}

func recordMetrics(accessInfo apiconnection.Info, conf Configuration) {
	go func() {
		client := api.NewClient(accessInfo, &http.Transport{
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
func ServeMetrics(secretsFilePath string, configurationFilePath string) {
	configuration, err := GetConfigurationFromFile(configurationFilePath)
	if err != nil {
		log.Fatal(err)
	}

	connectionInfo, err := apiconnection.ReadFile(secretsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	recordMetrics(*connectionInfo, configuration)

	http.Handle(configuration.MetricsPath, promhttp.Handler())
	http.ListenAndServe(":"+configuration.ExposedPort, nil)
}
