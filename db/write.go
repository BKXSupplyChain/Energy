package db

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
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
	gob.Register(&types.SocketInfo{})
	gob.Register(&types.Proposal{})
	gob.Register(&types.Contract{})
	gob.Register(&types.UserData{})

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

func TokenGetPower(id string) (power float32) {
	sensColl := session.DB(getMongoDatabase()).C(getMongoMetricCollection())
	sensQ := sensColl.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Sort("-Timestamp").Limit(2).Iter()
	var prevP, curP types.SensorPacket
	sensQ.Next(&curP)
	sensQ.Next(&prevP)
	if curP.Timestamp != prevP.Timestamp {
		return float32(curP.Total-prevP.Total) / float32(curP.Timestamp-prevP.Timestamp) * 1e9
	} else {
		return float32(math.Inf(1))
	}
}

func GetNewID() string {
	return string(bson.NewObjectId())
}

type generalPacket struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	B  []byte
}

func Add(obj interface{}, id string) error {
	var buff bytes.Buffer
	gob.NewEncoder(&buff).Encode(obj)
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	return coll.Insert(bson.M{"_id": id, "b": buff.Bytes()})
}

func Get(obj interface{}, id string) (err error) {
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	var encoded generalPacket
	err = coll.FindId(id).One(&encoded)
	if err != nil {
		return
	}
	return gob.NewDecoder(bytes.NewBuffer(encoded.B)).Decode(&obj)
}

func Upsert(obj interface{}, id string) {
	var buff bytes.Buffer
	gob.NewEncoder(&buff).Encode(obj)
	coll := session.DB(getMongoDatabase()).C(getMongoGeneralCollection())
	coll.UpsertId(id, bson.M{"$set": bson.M{"_id": id, "b": buff.Bytes()}})
}
