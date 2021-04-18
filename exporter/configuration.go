package exporter

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Configuration Overall configuration including exposed ports and so on
type Configuration struct {
	ExposedPort string `yaml:"exposed_port"`
	QueryPeriod int    `yaml:"query_period_sec"`
	MetricsPath string `yaml:"metrics_path"`
}

func (c *Configuration) parseYaml(rawContent []byte) error {
	err := yaml.Unmarshal(rawContent, &c)
	if err != nil {
		return err
	}

	return nil
}

//GetConfigurationFromFile Get connection information from a YAML file
func GetConfigurationFromFile(filePath string) (*Configuration, error) {
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var c Configuration

	err2 := c.parseYaml(rawContent)
	if err2 != nil {
		return nil, err
	}

	return &c, nil
}
