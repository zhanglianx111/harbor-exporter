package harbor

import "github.com/zhanglianx111/harbor-exporter/pkgs/config"

const (
	LOGIN     	= "/c/login"
	HEALTH    	= "/api/health"
	LDAP_PING 	= "/api/ldap/ping"
	USERS		= "/api/users"
	VOLUME		= "/api/systeminfo/volumes"

)

var Config *config.Conf

func init() { // read user、passwd and url of harbor
	Config = config.GetConfig()
	if Config == nil {
		return
	}
}
