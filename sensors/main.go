package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/BKXSupplyChain/Energy/db"
	"github.com/BKXSupplyChain/Energy/sensors/conf"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
	"github.com/globalsign/mgo"
)

var mgoSession *mgo.Session

func recieveData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reqBody)
	var pck types.SensorPacket
	err = json.Unmarshal(reqBody, &pck)
	if err != nil {
		utils.CheckNotFatal(err)
		return
	}
	pck.Timestamp = time.Now().Unix()
	db.WriteSensorsPacket(mgoSession, &pck)
	fmt.Fprintf(w, "OK") // send data to client side
}

func main() {
	log.Println("Starting sensors")
	conf.LoadConfig("config.json")
	var err error
	log.Println("Config loaded")
	mgoSession, err = mgo.Dial("mongodb://mongo:27017") // MUST BE MOVED TO DB
	log.Println(conf.GetMongoConnectionString()) // IS EMPTY! REPLACE CONFIG SYSTEM!
	//mgoSession, err = mgo.Dial(conf.GetMongoConnectionString())
	utils.CheckFatal(err)
	log.Println("DB connected")
	defer mgoSession.Close()

	http.HandleFunc("/", recieveData)       // set router
	err = http.ListenAndServe(":1000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
