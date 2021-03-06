package proposalRest

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"

    "github.com/BKXSupplyChain/Energy/types"
    "github.com/BKXSupplyChain/Energy/db"
    //"bytes"
    "fmt"
	//"errors"
    //"io/ioutil"
    "os/exec"
    "crypto/md5"
    "encoding/hex"
)

var proposals []types.Proposal


func GetSocketHash(From string, To string) string {
    concat := fmt.Sprintf("%s%s", From, To)
    hasher := md5.New()
    hasher.Write([]byte(concat))
    socketHash := hex.EncodeToString(hasher.Sum(nil))
    socketHash = socketHash[:12]
    return socketHash
}

//---------------------------------
//     REST METHODS
//---------------------------------

func CreateProposal(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var prop types.Proposal
    _ = json.NewDecoder(r.Body).Decode(&prop)
    prop.ID = params["ID"]
    proposals = append(proposals, prop)
    json.NewEncoder(w).Encode(proposals)

    socketID := GetSocketHash(prop.From, prop.To)
    var socket types.SocketInfo
    db.Get(&socket, socketID)
    socket.proposals = append(socket.proposals, prop)
    db.Upsert(&socket, socketID)
}

func DeleteProposal(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range proposals {
        if item.ID == params["ID"] {
            proposals = append(proposals[:index], proposals[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(proposals)
    }
}


func GetAll(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(proposals)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range proposals {
        if item.ID == params["ID"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&types.Proposal{})
}

func UpdateProposal(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range proposals{
        if item.ID == params["ID"] {
            proposals = append(proposals[:index], proposals[index+1:]...)
            var proposal types.Proposal
            _ = json.NewDecoder(r.Body).Decode(&proposal)
            proposal.ID = params["ID"]
            proposals = append(proposals, proposal) 
            json.NewEncoder(w).Encode(proposal)
            return
        }
    }
    json.NewEncoder(w).Encode(proposals)
}

//--------------------------------------


func Init() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/proposals", GetAll).Methods("GET")
    router.HandleFunc("/proposals/{ID}", GetOne).Methods("GET")
    router.HandleFunc("/proposals/{ID}", CreateProposal).Methods("POST")
    router.HandleFunc("/proposals/{ID}", DeleteProposal).Methods("DELETE")
    router.HandleFunc("/proposals/{ID}", UpdateProposal).Methods("PUT")

    srv := &http.Server{
        Handler:      router,
        Addr:         "localhost:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())

    fmt.Printf("REST server is set up.")
    return router
}

