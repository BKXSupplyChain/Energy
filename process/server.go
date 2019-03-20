package server
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/http"
    "net"
    "net/rpc"
)

func main() {
    cp := new(types.CertificatedPair) //check here 
    err := rpc.register(cp)
    if (err != nil) {
         log.Fatalf("Error of format of CertificatedPair: %s", err)
    }
    rpc.HandleHTTP()
    listener,err = net.Listen("tcp", ":8080")
    if (err != nil) {
          log.Fatalf("Can't listen on port 8080: %s", err)
    }
    err = http.Serve(listener, nil)
    if (err != nil) {
        log.Fatalf("Error serving: %s", err)
    }
}
