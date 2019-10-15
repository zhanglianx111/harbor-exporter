package harbor

import (
	"fmt"
	"encoding/json"
	"github.com/zhanglianx111/harbor-exporter/pkgs/config"
)


type VolumeInfo struct {
	Total uint64 `json: "total"`
	Free  uint64 `json: "free"`
}

// SystemInfo models for system info.
type SystemInfo struct {
	HarborStorage Storage `json:"storage"`
}

// Storage models for storage.
type Storage struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

func GetVolumeInfo() map[string]uint64 {
	volumeInfo := map[string]uint64{}

	cfg := config.GetConfig()
	if cfg.Harbor == "" {
		return volumeInfo
	}

	volume := SystemInfo{Storage{}}
	volumeUrl := "http://" + cfg.Harbor + VOLUME
	bodyByte := get(volumeUrl)
	//fmt.Printf("volume from harbor: %v\n", string(bodyByte))
	err := json.Unmarshal(bodyByte, &volume)
	if err != nil {
		fmt.Println(err.Error())
		return volumeInfo
	}

	// volume infomations
	volumeInfo["total"] = volume.HarborStorage.Total
	volumeInfo["free"] = volume.HarborStorage.Free

	return volumeInfo
}