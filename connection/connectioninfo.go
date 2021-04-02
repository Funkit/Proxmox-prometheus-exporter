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

type Info struct {
	Address  string   `yaml:"apiaddress"`
	UserId   UserID   `yaml:"userid"`
	ApiToken ApiToken `yaml:"apitoken"`
}

func (c *Info) ReadFile(filePath string) *Info {
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
