package exporter

import "github.com/prometheus/client_golang/prometheus"

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
	maxCpu = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "PVE",
			Subsystem: "hardware",
			Name:      "cpu_max",
			Help:      "Maximum number of allocated CPU.",
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
	ramLoad = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "PVE",
			Subsystem: "hardware",
			Name:      "ram_current",
			Help:      "Current RAM usage",
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
	maxRam = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "PVE",
			Subsystem: "hardware",
			Name:      "ram_max",
			Help:      "Maximum RAM allocated",
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
