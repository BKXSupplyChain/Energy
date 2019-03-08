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

func GetMongoSensorsDatabase() string {
	return config.Get("mongo", "sensors").String("")
}

func GetMongoMetricCollection() string {
	return config.Get("mongo", "metric").String("")
}

func GetMongoConnectionString() string {
	return config.Get("mongo", "connection_string").String("")
}
