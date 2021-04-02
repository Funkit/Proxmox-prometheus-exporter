package api

// /nodes
type Node struct {
	SslFingerprint string `json:"ssl_fingerprint,omitempty"`
	Status         string `json:"status"`
	Node           string `json:"node,omitempty"`
}

type Nodes struct {
	NodeList []Node `json:"data"`
}

// /cluster/resources
type VMResource struct {
	Name              string  `json:"name"`
	VMID              int     `json:"vmid"`
	Pool              string  `json:"pool,omitempty"`
	Node              string  `json:"node"`
	Status            string  `json:"status,omitempty"`
	Uptime            int     `json:"uptime"`
	AllocatedCPUCores int     `json:"maxcpu"`
	CPU               float64 `json:"cpu"`    // %
	AllocatedRAMBytes int     `json:"maxmem"` // in bytes
	RAM               int     `json:"mem"`    // in bytes
	Template          int     `json:"template"`
}

type NodeResource struct {
	Node   string `json:"node"`
	Status string `json:"status"`
}
