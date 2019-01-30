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

func GetServerIP(idx string) []byte {
	IPStr := config.Get("servers", idx, "ip").String("")
	return utils.ConvertIPToBytes(IPStr)
}

func GetServerPort(idx string) int {
	return config.Get("servers", idx, "port").Int(0)
}
