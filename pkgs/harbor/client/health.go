package harbor

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
)

const (
	Redis       = "redis"
	Core        = "core"
	Registryctl = "registryctl"
	Database    = "database"
	Portal      = "portal"
	Jobservice  = "jobservice"
	Registry    = "registry"
	Harbor      = "harbor"
)

type HealthComponent struct {
	Name   string `json: "name"`
	Status string `json: "status"`
}

type Health struct {
	Status     string            `json: "status"`
	Components []HealthComponent `json: "components"`
}

/*
	-1: 内部错误
	0: 整个harbor不健康
	1: 整个harbor健康
*/
func GetHealthStatus() map[string]int8 {
	status := map[string]int8{}

	cfg := config.GetConfig()
	if cfg.Harbor == "" {
		return status
	}

	healths := &Health{}
	healthUrl := cfg.Harbor + HEALTH
	bodyByte := get(healthUrl)
	if len(bodyByte) == 0 {
		log.Warn("get response body is nil")
		return status
	}
	err := json.Unmarshal(bodyByte, healths)
	if err != nil {
		log.Errorf("json unmarshal error: %v\n",err.Error())
		return status
	}

	// harbor status
	if healths.Status == "healthy" {
		status[Harbor] = 1
	} else {
		status[Harbor] = 0
	}
	// redis status
	for _, comp := range healths.Components {
		if comp.Status == "healthy" {
			status[comp.Name] = 1
		} else {
			status[comp.Name] = 0
		}
	}

	return status
}
