package main

import (
	"github.com/BKXSupplyChain/Energy/backend/conf"
	"github.com/BKXSupplyChain/Energy/backend/web"
	"github.com/BKXSupplyChain/Energy/db"
	"github.com/BKXSupplyChain/Energy/utils"
	"log"
)

func main() {
	log.Println("Starting backend")
	conf.LoadConfig("conf.json")
	var err error
	log.Println("Config loaded")
	err = db.DBCreateSession()
	utils.CheckFatal(err)
	defer db.Close()

	web.Serve()
}
