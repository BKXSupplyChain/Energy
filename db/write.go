package db

import (
	"github.com/BKXSupplyChain/Energy/sensors/conf"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
	"github.com/globalsign/mgo"
)

func WriteSensorsPacket(session *mgo.Session, pck *types.SensorPacket) {
	coll := session.DB(conf.GetMongoSensorsDatabase()).C(conf.GetMongoMetricCollection())
	err := coll.Insert(&pck)
	utils.CheckFatal(err)
}
