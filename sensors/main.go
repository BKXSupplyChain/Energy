package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reqBody)
	fmt.Fprintf(w, "OK") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":1000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
