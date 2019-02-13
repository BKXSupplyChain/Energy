package conf

import (
	"github.com/BKXSupplyChain/Energy/utils"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func LoadConfig(path string) {
	config.Load(file.NewSource(
		file.WithPath(path),
	))
}

func GetServerIPAsString(idx string) string {
	return config.Get("servers", idx, "ip").String("")
}

func GetServerIP(idx string) []byte {
	IPStr := GetServerIPAsString(idx)
	return utils.ConvertIPToBytes(IPStr)
}

func GetServerPort(idx string) int {
	return config.Get("servers", idx, "port").Int(0)
}

func GetServerPortAsString(idx string) string {
	return config.Get("servers", idx, "port").String("")
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
