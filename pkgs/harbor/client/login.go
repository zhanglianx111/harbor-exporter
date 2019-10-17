package harbor

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
	"os"
)

// c/login
// curl -X POST --header 'Content-Type: application/x-www-form-urlencoded;param=value' 'http://exporter.harbor.com/c/login' -i -k -d "principal=admin&password=Harbor12345"
func Login() {
	// get config for harbor
	/*
	cfg := config.GetConfig()
	if cfg == nil {
		return
	}
	*/
	user, password := config.Config.Username, config.Config.Password
	loginUrl := config.Config.Harbor + LOGIN
	log.Debugf("login from url:%s\n", loginUrl)

	cookie := postLogin(loginUrl, user, password)
	if cookie != "" {
		//log.Infof("get cookie is: %v\n", cookie)
		os.Setenv("cookie", cookie)
	} else {
		log.Errorf("no get cookie for user: %s\n", config.Config.Username)
	}

	if len(config.Config.Harbor) != 0 {
		os.Setenv("harbor", config.Config.Harbor)
	}
}
