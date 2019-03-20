package client
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/rpc"
    //TODO import backend
)


func main() {
    client, err := rpc.DialHTTP("tcp", ":8080")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    cp := .CreateCertificatedPair //TODO name of function in start which creates pair
    err = client.Call("CertificatedPair."/*TODO name of function of checking if all is correct*/,cp, {})
    if (err != nil) {
        log.Fatalf("Error in CertificatedPair: %s", err)
    }
}
