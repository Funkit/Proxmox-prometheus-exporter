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

type UserID struct {
	Username string `yaml:"username"`
	IdRealm  string `yaml:"idrealm"`
}

type ApiToken struct {
	Id    string `yaml:"id"`
	Token string `yaml:"token"`
}

type Info struct {
	Address  string   `yaml:"apiaddress"`
	UserId   UserID   `yaml:"userid"`
	ApiToken ApiToken `yaml:"apitoken"`
}

func (c *Info) parseYaml(rawContent []byte) error {
	err := yaml.Unmarshal(rawContent, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Info) ReadFile(filePath string) error {
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return c.parseYaml(rawContent)
}
