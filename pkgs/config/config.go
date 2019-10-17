package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Username string `yaml: "username"`
	Password string `yaml: "password"`
	Harbor   string `yaml: "harbor"`
}

func GetConfig() *Conf {
	c := new(Conf)
	yamlFile, err := ioutil.ReadFile("/etc/harbor-exporter/config.yaml")
	if err != nil {
		log.Errorf("read harbor config error: %v\n", err.Error())
		return nil
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Errorf("unmarshal config file err: %v\n", err.Error())
		return nil
	}

	return c
}
