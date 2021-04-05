package api

import "encoding/json"

// /nodes
type Node struct {
  Name           string `json:"node,omitempty"`
  Status         string `json:"status"`
  SslFingerprint string `json:"ssl_fingerprint,omitempty"`
}

func (node *Node) ParseMap(element map[string]interface{}) error {
  jsonbody, err := json.Marshal(element)
  if err != nil {
    return err
  }

  if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
    return err
  }
  return nil
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

func (node *VMResource) ParseMap(element map[string]interface{}) error {
  jsonbody, err := json.Marshal(element)
  if err != nil {
    return err
  }

  if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
    return err
  }
  return nil
}

type NodeResource struct {
  Node   string `json:"node"`
  Status string `json:"status"`
}

func (node *NodeResource) ParseMap(element map[string]interface{}) error {
  jsonbody, err := json.Marshal(element)
  if err != nil {
    return err
  }

  if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
    return err
  }
  return nil
}

// /nodes/<node name>/network

type NodeNetworkInterface struct {
  Name          string   `json:"iface"`
  InterfaceType string   `json:"type"`
  Active        int      `json:"active"`
  IPAddress     string   `json:"address"`
  Gateway       string   `json:"gateway"`
  Autostart     int      `json:"autostart"`
  BridgePorts   string   `json:"bridge_ports"`
  CIDR          string   `json:"cidr"`
  Families      []string `json:"families"`
  Options       []string `json:"options"`
}
