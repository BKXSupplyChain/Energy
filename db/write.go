package db

import (
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
	"github.com/globalsign/mgo"
	"github.com/micro/go-config"
	"log"
)

var session *mgo.Session

func getMongoSensorsDatabase() string {
	return config.Get("mongo", "sensors").String("")
}

func getMongoMetricCollection() string {
	return config.Get("mongo", "metric").String("")
}

func DBCreateSession() (err error) {
	addr := config.Get("mongo", "connection_string").String("")
	log.Println("connecting to ", addr)
	session, err = mgo.Dial(addr)
	if err != nil {
		return
	}
	log.Println("DB connected")
	return
}

func Close() {
	session.Close()
}

func WriteSensorsPacket(pck *types.SensorPacket) {
	coll := session.DB(getMongoSensorsDatabase()).C(getMongoMetricCollection())
	err := coll.Insert(&pck)
	utils.CheckFatal(err)
}
