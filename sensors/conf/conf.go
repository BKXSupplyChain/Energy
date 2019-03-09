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
