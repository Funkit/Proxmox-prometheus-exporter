package connectioninfo

// Example yml file //
// ipaddr: 192.0.2.12
// userid:
//   username: toto
//   idrealm: pve
// apitoken:
//   id: prometheus
//   token: AAAAABBBBBCCCCCDDDDD

import (
	"io/ioutil"
	"log"

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

type ConnectionInfo struct {
	Address  string   `yaml:"ipaddr"`
	UserId   UserID   `yaml:"userid,omitempty"`
	ApiToken ApiToken `yaml:"apitoken,omitempty"`
}

func (c *ConnectionInfo) ReadFile(filePath string) *ConnectionInfo {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
