package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Conf struct {
	Username string `yaml: "username"`
	Password string `yaml: "password"`
	Harbor   string `yaml: "harbor"`
}

var Config *Conf

func init() {
	getUserInfo()
}

func getUserInfo() {
	Config = new(Conf)
	//get username
	user, err := ioutil.ReadFile("/etc/harbor/username")
	if err != nil {
		log.Errorf("read harbor config error: %v\n", err.Error())
		return
	}

	//get password
	password, err := ioutil.ReadFile("/etc/harbor/password")
	if err != nil {
		log.Errorf("read harbor config error: %v\n", err.Error())
		return
	}

	//get harbor address
	harbor, err := ioutil.ReadFile("/etc/harbor/harbor")
	if err != nil {
		log.Errorf("read harbor config error: %v\n", err.Error())
		return
	}

	log.Debugf("username: %s, password: %s, harbor: %s\n", user, password, harbor)

	if len(user) == 0 || len(password) == 0 || len(harbor) == 0 {
		return
	}

	Config.Username = fmt.Sprintf("%s", user)
	Config.Password = fmt.Sprintf("%s", password)
	Config.Harbor = fmt.Sprintf("%s", harbor)
}

func GetUsername() string {
	return Config.Username
}

func GetPassword() string {
	return Config.Password
}

func GetHarborHost() string {
	return Config.Harbor
}
