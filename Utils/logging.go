package utils

import (
	"log"
)

func CheckNotFatal(err error) {
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
