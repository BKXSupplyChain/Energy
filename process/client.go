package client
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/rpc"
)

func sendProposal(proposalMsg types.ProposalMessage, userPublicKey string, neighbourPublicKey string) {
    socketID := string(sha256.Sum256([]byte(userPublicKey+NeighbourPublicKey))[:12])
    var socket types.SocketInfo
    db.Get(&socket, socketID)
    client, err := rpc.DialHTTP("tcp", socket.NeighborAddr + ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("types.ProposalMsg.OnProposalReceived", proposalMsg, {})
    if (err != nil) {
        log.Fatalf("Error in ProposalMessage: %s", err)
    }
}

func sendCertificate(certif types.Certificate) {
    client, err := rpc.DialHTTP("tcp", /*TODO address*/ ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("types.Certificate."/*TODO name of function of checking if all is correct*/, certif, {})
    if (err != nil) {
        log.Fatalf("Error in CertificatedPair: %s", err)
    }
}
