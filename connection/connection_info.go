package connection

// Example yml file //
// address: 192.0.2.12
// userid:
//   username: toto
//   idrealm: pve
// apitoken:
//   id: prometheus
//   token: AAAAABBBBBCCCCCDDDDD

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//UserID User name and realm
type UserID struct {
	Username string `yaml:"username"`
	IDRealm  string `yaml:"idrealm"`
}

//APIToken Token ID and actual token
type APIToken struct {
	ID    string `yaml:"id"`
	Token string `yaml:"token"`
}

//Info token-based API access information
type Info struct {
	Address  string   `yaml:"apiaddress"`
	UserID   UserID   `yaml:"userid"`
	APIToken APIToken `yaml:"apitoken"`
}

func (c *Info) parseYaml(rawContent []byte) error {
	err := yaml.Unmarshal(rawContent, &c)
	if err != nil {
		return err
	}

	return nil
}

//GetInfoFromFile Get connection information from a YAML file
func GetInfoFromFile(filePath string) (*Info, error) {
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var c Info

	err2 := c.parseYaml(rawContent)
	if err2 != nil {
		return nil, err
	}

	return &c, nil
}
