package db

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/micro/go-config"
	"log"
	"math"
)

var session *mgo.Session

func getMongoDatabase() string {
	return config.Get("mongo", "sensors").String("sensors")
}

func getMongoMetricCollection() string {
	return config.Get("mongo", "metric").String("metric")
}

func getMongoGeneralCollection() string {
	return config.Get("mongo", "general").String("general")
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
	coll := session.DB(getMongoDatabase()).C(getMongoMetricCollection())
	err := coll.Insert(&pck)
	utils.CheckFatal(err)
}

//======== BACKEND ONLY ============

func getMongoUsersCollection() string {
	return config.Get("mongo", "users").String("users")
}

func b64ToObjectId(in string) bson.ObjectId {
	byteID, _ := base64.URLEncoding.DecodeString(in)
	return bson.ObjectId(byteID)
}

func objectIdToB64(in bson.ObjectId) string {
	bytedID, _ := hex.DecodeString(in.String())
	return base64.URLEncoding.EncodeToString(bytedID)
}

func ForgeToken(d *types.SocketInfo) string {
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	var objectId bson.ObjectId = bson.NewObjectId()
	utils.CheckFatal(coll.Insert(bson.M{"_id": objectId}, &d))
	return string(objectId)
}

func TokenGetPower(id string, fromTime int64) (power float32) {
	sensColl := session.DB(getMongoDatabase()).C(getMongoMetricCollection())
	sensQ := sensColl.Find(bson.M{"sensorID": bson.ObjectIdHex(id), "Timestamp": bson.M{"$gt": fromTime * 1e6}}).Sort("-Timestamp").Limit(2).Iter()
	var prevP, curP types.SensorPacket
	if !sensQ.Next(&curP) || !sensQ.Next(&prevP) {
		return 0
	}
	if curP.Timestamp != prevP.Timestamp {
		return float32(curP.Total-prevP.Total) / float32(curP.Timestamp-prevP.Timestamp) * 1e9
	} else {
		return float32(math.Inf(1))
	}
}

func GetNewID() string {
	return string(bson.NewObjectId())
}

func strToId(a string) bson.ObjectId {
	hashed := md5.Sum([]byte(a))
	return bson.ObjectId(hashed[:12])
}

func Add(obj interface{}, id string) error {
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	err := coll.Insert(bson.M{"_id": strToId(id)})
	if err != nil {
		log.Println("ADD ERROR:", coll.Name, "->", id, obj, ":", err)
		return err
	}
	Upsert(obj, id)
	return nil
}

func Get(obj interface{}, id string) (err error) {
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	return coll.FindId(strToId(id)).One(obj)
}

func Upsert(obj interface{}, id string) {
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	coll.UpsertId(strToId(id), bson.M{"$set": obj})
}

func getEnergy(SocketID string) uint64 {
	sensColl := session.DB(getMongoDatabase()).C(getMongoMetricCollection())
	var pac types.SensorPacket
	sensColl.Find(bson.M{"sensorID": bson.ObjectIdHex(SocketID)}).Sort("-Timestamp").One(&pac)
	return pac.Total
}
