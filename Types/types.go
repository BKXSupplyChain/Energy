package types

import (
	"net"

	"github.com/globalsign/mgo/bson"
)

type SensorListener struct {
	IP        []byte
	Port      int
	SensorsID []bson.ObjectId
	TCPL      *net.TCPListener
}
