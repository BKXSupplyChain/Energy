package db

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/micro/go-config"
	"log"
	"math"
	"math/big"
)

type TokenId bson.ObjectId

var session *mgo.Session

func getMongoDatabase() string {
	return config.Get("mongo", "sensors").String("sensors")
}

func getMongoMetricCollection() string {
	return config.Get("mongo", "metric").String("metric")
}

func getMongoSocketsCollection() string {
	return config.Get("mongo", "sockets").String("sockets")
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
	coll := session.DB(getMongoDatabase()).C(getMongoSocketsCollection())
	var objectId bson.ObjectId = bson.NewObjectId()
	utils.CheckFatal(coll.Insert(bson.M{"_id": objectId}, &d))
	return objectId.Hex()
}

func TokenGetPower(t string) (power float32) {
	sensColl := session.DB(getMongoDatabase()).C(getMongoMetricCollection())
	sensQ := sensColl.Find(bson.M{"_id": bson.ObjectIdHex(t)}).Sort("-Timestamp").Limit(2).Iter()
	var prevP, curP types.SensorPacket
	sensQ.Next(&curP)
	sensQ.Next(&prevP)
	if curP.Timestamp != prevP.Timestamp {
		return float32(curP.Total-prevP.Total) / float32(curP.Timestamp-prevP.Timestamp) * 1e9
	} else {
		return float32(math.Inf(1))
	}
}

func addUser(u *types.UserData) {
	coll := session.DB(getMongoDatabase()).C(getMongoUsersCollection())
	err := coll.Insert(&u)
	utils.CheckFatal(err)
}

func GetNewID() string {
	return string(bson.NewObjectId())
}

func GetSocket(id string) types.SocketInfo {
}

func UpsertSocket(socInfo *types.SocketInfo, id string) { ///Update + (Insert if id missing)
}

func GetProposal(id string) types.Proposal {
}

func InsertProposal(socInfo *types.Proposal, id string) {
}

func ApplySertificate(contract string, amount big.Int, signature []byte) { ///Check everything yourself
}

func GetContract(id string) types.Contract {
}

func PutContract(con *types.Contract, id string) error { ///will also set in socket
	return errors.New("ID Already Used") ///Could throw this error. Treat this as malicious neighbour
}
