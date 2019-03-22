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

func AddInDatabase(proposal types.Proposal)  {
    db.Add(proposal, proposal.ID) //TODO ask about ID
}

func updateSocket(proposal types.Proposal, userPublicKey /*TODO type */, NeighbourPublicKey /*TODO type*/) {
    socketID := string(sha256.Sum256([]byte(userPublicKey+NeighbourPublicKey))[:12])
    var socket = types.SocketInfo
    db.Get(&socket, socketID)
    socket.proposals = append(socket.proposals, /* TODO what to add */
    db.Upsert(&socket, socketID) //TODO maybe implement Update
}

func (t *types.Proposal) OnProposalReceived(proposal types.Proposal) {
    putInDb(proposal)
    updateSocket(proposal, /*TODO key*/, /*TODO key*/)
}

func main() {
    proposal := new(types.Proposal)
    err := rpc.register(proposal)
    if (err != nil) {
         log.Fatalf("Error of format of proposal: %s", err)
    }
    rpc.HandleHTTP()
    listener,err = net.Listen("tcp", ":40")
    if (err != nil) {
          log.Fatalf("Can't listen on port 40: %s", err)
    }
    err = http.Serve(listener, nil)
    if (err != nil) {
        log.Fatalf("Error serving: %s", err)
    }
}
