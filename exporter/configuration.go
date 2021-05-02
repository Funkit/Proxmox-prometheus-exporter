package exporter

import (
	"github.com/Funkit/proxmox-prometheus-exporter/common"
	"gopkg.in/yaml.v2"
)

//Configuration Overall configuration including exposed ports and so on
type Configuration struct {
	ExposedPort string `yaml:"exposed_port"`
	QueryPeriod int    `yaml:"query_period_sec"`
	MetricsPath string `yaml:"metrics_path"`
	SecretsPath string `yaml:"secrets_file_path"`
}

func (c *Configuration) parseYaml(rawContent []byte) error {
	err := yaml.Unmarshal(rawContent, &c)
	if err != nil {
		return err
	}

	return nil
}

//GetConfigurationFromFile Get connection information from a YAML file
func GetConfigurationFromFile(filePath string) (Configuration, error) {
	var c Configuration

	if err := common.GetInfo(filePath, &c); err != nil {
		panic(err)
	}

	return c, nil
}
