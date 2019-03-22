package conf

import (
	"github.com/BKXSupplyChain/Energy/utils"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func LoadConfig(path string) {
	err := config.Load(file.NewSource(
		file.WithPath(path),
	))
	utils.CheckFatal(err)
}

func GetSelfAddress() string {
	return config.Get("selfAddress").String("") /// TODO: get it automaticly
}
