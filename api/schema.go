package api

type Node struct {
	SslFingerprint string `json:"ssl_fingerprint,omitempty"`
	Status         string `json:"status,omitempty"`
	Node           string `json:"node,omitempty"`
}

type Nodes struct {
	NodeList []Node `json:"data"`
}
