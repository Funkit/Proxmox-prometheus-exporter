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
