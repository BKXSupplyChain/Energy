package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// The ProcessMessage Type
type ProcessMessage struct {
    IP        string   `json:"ip,omitempty"`
    TotalCost int32    `json:"TotalCost,omitempty"`
    Salt      int32    `json:"Salt,omitempty"`
    TTL       int32    `json:"TTL,omitempty"`
}

var messages []ProcessMessage

// Display all from the messages var
func GetMessages(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(messages)
}

// Display a single data
func GetMessage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range messages {
        if item.IP == params["ip"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&ProcessMessage{})
}

// create a new item
func CreateMessage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var msg ProcessMessage
    _ = json.NewDecoder(r.Body).Decode(&msg)
    msg.IP = params["ip"]
    messages = append(messages, msg)
    json.NewEncoder(w).Encode(messages)
}

// Delete an item
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range messages {
        if item.IP == params["ip"] {
            messages = append(messages[:index], messages[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(messages)
    }
}

func CallFinishSup() {

}

func CallFinishExp() {

}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    messages = append(messages, ProcessMessage{IP: "1.0.0.0", TotalCost: 173000, Salt: 3, TTL: 4})
    messages = append(messages, ProcessMessage{IP: "2.0.0.0", TotalCost: 174000, Salt: 3, TTL: 4})
    router.HandleFunc("/messages", GetMessages).Methods("GET")
    router.HandleFunc("/messages/{ip}", GetMessage).Methods("GET")
    router.HandleFunc("/messages/{ip}", CreateMessage).Methods("POST")
    router.HandleFunc("/messages/{ip}", DeleteMessage).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
