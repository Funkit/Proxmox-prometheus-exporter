package api

import (
	"encoding/json"
)

type Results struct {
	Rows []map[string]interface{} `json:"data"`
}

func ParseClusterResources(responseBody []byte) ([]NodeResource, []VMResource, error) {
	var buffer Results

	if err := json.Unmarshal([]byte(responseBody), &buffer); err != nil {
		return nil, nil, err
	}

	var nodeList []NodeResource
	var vmList []VMResource
	for _, row := range buffer.Rows {
		if row["type"] == "node" {
			jsonbody, err := json.Marshal(row)
			if err != nil {
				return nil, nil, err
			}

			var buffer NodeResource
			if err := json.Unmarshal([]byte(jsonbody), &buffer); err != nil {
				return nil, nil, err
			}

			nodeList = append(nodeList, buffer)
		}
		if row["type"] == "qemu" {
			jsonbody, err := json.Marshal(row)
			if err != nil {
				return nil, nil, err
			}

			var buffer VMResource
			if err := json.Unmarshal([]byte(jsonbody), &buffer); err != nil {
				return nil, nil, err
			}

			vmList = append(vmList, buffer)
		}
	}
	return nodeList, vmList, nil
}

func ParseNodes(responseBody []byte) ([]Node, error) {
	var buffer Results

	if err := json.Unmarshal([]byte(responseBody), &buffer); err != nil {
		return nil, err
	}

	var nodeList []Node
	for _, row := range buffer.Rows {
		jsonbody, err := json.Marshal(row)
		if err != nil {
			return nil, err
		}

		var buffer Node
		if err := json.Unmarshal([]byte(jsonbody), &buffer); err != nil {
			return nil, err
		}

		nodeList = append(nodeList, buffer)
	}
	return nodeList, nil
}
