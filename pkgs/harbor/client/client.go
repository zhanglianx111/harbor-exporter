package harbor

import (
	"encoding/base64"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
	"os"
)

const (
	LOGIN     	= "/c/login"
	HEALTH    	= "/api/health"
	LDAP_PING 	= "/api/ldap/ping"
	USERS		= "/api/users"
	VOLUME		= "/api/systeminfo/volumes"
    STATISTICS	= "/api/statistics"

)

func init() { // read user、passwd and url of harbor
	/*
	Config = config.GetConfig()
	if Config == nil {
		return
	}
	*/
	// baisc auth

	tok := config.Config.Username+":"+config.Config.Password
	hashStr := base64.StdEncoding.EncodeToString([]byte(tok))
	if len(hashStr) != 0 {
		os.Setenv("baseAuth", hashStr)
	}
}


