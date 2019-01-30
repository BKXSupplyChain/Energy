package main

import (
	"log"
	"net"
	"os"

	"github.com/BKXSupplyChain/Energy/sensors/conf"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/utils"
)

// Can be expended in many listeners
func startListening(idx string) {
	sl := types.SensorListener{IP: conf.GetServerIP(idx), Port: conf.GetServerPort(idx)}
	tcpl, err := net.ListenTCP("tcp", &net.TCPAddr{IP: sl.IP, Port: sl.Port, Zone: ""})
	utils.CheckFatal(err)
	sl.TCPL = tcpl
	log.Printf("Listener %s is running\n", idx)
	for {
		conn, err := sl.TCPL.Accept()
		if err != nil {
			utils.CheckNotFatal(err)
			continue
		}
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	log.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		len, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return
		}
		log.Println("Number of bytes: ", len)
		log.Println(string(buf[0:]))
	}
}

func main() {
	// Set up logging
	file, err := os.OpenFile("main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	utils.CheckFatal(err)
	defer file.Close()
	log.SetOutput(file)
	startListening("1")
}
