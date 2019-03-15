package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"

    "./types"
    "bytes"
)

var proposals []types.Proposal


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

func SendProposal(P types.Proposal) {
    proposals = append(proposals, P)
    //json.NewEncoder(w).Encode(proposals)
}

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

    return router
}

func main() {
    Init()
    newprop := types.Proposal{ID: "1.0.0.3",
                                                 Supplier : "A",
                                                 Consumer : "B", 
                                                 Price: 2, 
                                                 TotalCost : 200001,
                                                 DeltaPrice : 20, 
                                                 Salt: 2, 
                                                 TTL : 7}
    SendProposal(newprop)
}

