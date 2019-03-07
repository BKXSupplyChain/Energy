package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"

    "./types"
)

var proposals []types.Proposal

func CreateProposal(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var prop types.Proposal
    _ = json.NewDecoder(r.Body).Decode(&prop)
    prop.IP = params["IP"]
    proposals = append(proposals, prop)
    json.NewEncoder(w).Encode(proposals)
}

func DeleteProposal(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range proposals {
        if item.IP == params["IP"] {
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
        if item.IP == params["IP"] {
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
        if item.IP == params["IP"] {
            proposals = append(proposals[:index], proposals[index+1:]...)
            var proposal types.Proposal
            _ = json.NewDecoder(r.Body).Decode(&proposal)
            proposal.IP = params["IP"]
            proposals = append(proposals, proposal) 
            json.NewEncoder(w).Encode(proposal)
            return
        }
    }
    json.NewEncoder(w).Encode(proposals)
}

func main() {
    router := mux.NewRouter()
    proposals = append(proposals, types.Proposal{IP: "1.0.0.0", Price: 1, TotalCost: 200000, DeltaPrice: 10, Salt: 42, TTL: 5})
    router.HandleFunc("/proposals", GetAll).Methods("GET")
    router.HandleFunc("/proposals/{IP}", GetOne).Methods("GET")
    router.HandleFunc("/proposals/{IP}", CreateProposal).Methods("POST")
    router.HandleFunc("/proposals/{IP}", DeleteProposal).Methods("DELETE")
    router.HandleFunc("/proposals/{IP}", UpdateProposal).Methods("PUT")
    log.Fatal(http.ListenAndServe(":8000", router))
}

