package config

import (
	"fmt"
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
	yamlFile, err := ioutil.ReadFile("/Users/zhanglx/go/src/github.com/zhanglianx111/harbor-exporter/config/config.yaml")
	if err != nil {
		fmt.Printf("read harbor config error: %v\n", err.Error())
		return nil
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("unmarshal config file err: %v\n", err.Error())
		return nil
	}

	return c
}
