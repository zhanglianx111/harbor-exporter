package harbor

import (
	"encoding/json"
	"fmt"
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
	healthUrl := "http://" + cfg.Harbor + HEALTH
	bodyByte := get(healthUrl)

	err := json.Unmarshal(bodyByte, healths)
	if err != nil {
		fmt.Println(err.Error())
		return status
	}
	fmt.Printf("%v\n", healths.Components)
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
