package harbor

import (
	"fmt"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
	"os"
)

// c/login
// curl -X POST --header 'Content-Type: application/x-www-form-urlencoded;param=value' 'http://exporter.harbor.com/c/login' -i -k -d "principal=admin&password=Harbor12345"
func login() {
	// get config for harbor
	cfg := config.GetConfig()
	if cfg == nil {
		fmt.Println("1")
		return
	}
	user, passwd := cfg.Username, cfg.Password
	loginUrl := cfg.Harbor + LOGIN
	fmt.Printf("login ->url:%s username: %s, password: %s\n", loginUrl, user, passwd)
	cookie := postLogin(loginUrl, cfg.Username, cfg.Password)
	fmt.Printf("login cookio: %v\n", cookie)

	if cookie != "" {
		os.Setenv("cookie", cookie)
	}
}
