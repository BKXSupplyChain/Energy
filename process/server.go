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
    db.Add(proposal, proposal.ID)
}

func UpdateSocket(proposal types.Proposal, userPublicKey string, NeighbourPublicKey string) {
    socketID := string(sha256.Sum256([]byte(userPublicKey+NeighbourPublicKey))[:12])
    var socket = types.SocketInfo
    db.Get(&socket, socketID)
    socket.proposals = append(socket.proposals, proposal.ID)
    db.Upsert(&socket, socketID)
}

func (t *types.ProposalMessage) OnProposalReceived(proposalMsg types.ProposalMessage) {
    AddInDatabase(proposalMsg.Proposal)
    UpdateSocket(proposal, proposalMsg.UserKey, /*TODO key*/)
}

func main() {
    cp := new(types.Certificate) 
    err := rpc.register(cp)
    if (err != nil) {
         log.Fatalf("Error of format of Certificate: %s", err)
    }
    proposalMsg := new(types.ProposalMsg)
    err := rpc.register(proposalMsg)
    if (err != nil) {
         log.Fatalf("Error of format of proposalMsg: %s", err)
    }
    rpc.HandleHTTP()
    listener,err = net.Listen("tcp", ":30")
    if (err != nil) {
          log.Fatalf("Can't listen on port 30: %s", err)
    }
    err = http.Serve(listener, nil)
    if (err != nil) {
        log.Fatalf("Error serving: %s", err)
    }
}
