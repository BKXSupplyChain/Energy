package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/BKXSupplyChain/Energy/sensors/conf"
	"github.com/BKXSupplyChain/Energy/utils"
)

func isExistsInDB(token string) bool {
	return token == "abcde"
}

// Can be expended in many listeners
func startListening(idx string) {
	connStr := conf.GetServerIPAsString(idx) + ":" + conf.GetServerPortAsString(idx)
	fmt.Println(connStr)
	listener, err := net.Listen("tcp", conf.GetServerIPAsString(idx)+":"+conf.GetServerPortAsString(idx))
	utils.CheckFatal(err)
	log.Printf("Listener %s is running\n", idx)
	for {
		conn, err := listener.Accept()
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
			fmt.Println("Closing connection...")
			log.Printf("Error: %s\n", err.Error())
			return
		}
		log.Println("Number of bytes: ", len)
		log.Println(string(buf[0:]))
		if isExistsInDB(string(buf[0:])) {
			conn.Write([]byte{1})
		} else {
			conn.Write([]byte{0})
		}
	}
}

func main() {
	// Set up logging
	file, err := os.OpenFile("main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	utils.CheckFatal(err)
	defer file.Close()
	log.SetOutput(file)
	// Set up config
	conf.LoadConfig("config.json")
	startListening("1")
}
