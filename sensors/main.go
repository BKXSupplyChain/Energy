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
)

func recieveData(w http.ResponseWriter, r *http.Request) {
	log.Println("RECIEVED", r.Body)
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
	pck.Timestamp = time.Now().UnixNano()
	db.WriteSensorsPacket(&pck)
	fmt.Fprintf(w, "OK") // send data to client side
}

func main() {
	log.Println("Starting sensors")
	conf.LoadConfig("config.json")
	var err error
	log.Println("Config loaded")
	err = db.DBCreateSession()
	utils.CheckFatal(err)
	defer db.Close()

	http.HandleFunc("/", recieveData)       // set router
	err = http.ListenAndServe(":1000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
