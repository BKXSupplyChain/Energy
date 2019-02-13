package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/BKXSupplyChain/Energy/utils"
)

type sensorPacket struct {
	UserToken     string
	NeighborToken string
	Value         int64
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reqBody)
	var pck sensorPacket
	err = json.Unmarshal(reqBody, &pck)
	if err != nil {
		utils.CheckNotFatal(err)
		return
	}
	fmt.Println(pck.UserToken)
	fmt.Println(pck.NeighborToken)
	fmt.Println(pck.Value)
	fmt.Fprintf(w, "OK") // send data to client side
}

func main() {
	t := time.Now().Unix()
	fmt.Println(t)
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":1000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
