package client
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/rpc"
)

func sendProposal(proposal types.Proposal, /* socket */, /* userKey */) {
    client, err := rpc.DialHTTP("tcp", ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("types.Proposal.OnProposalReceived", proposal, {})
    if (err != nil) {
        log.Fatalf("Error in Proposal: %s", err)
    }
}

func sendCertificate(certif types.Certificate) {
    client, err := rpc.DialHTTP("tcp", ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("types.Certificate."/*TODO name of function of checking if all is correct*/, certif, {})
    if (err != nil) {
        log.Fatalf("Error in CertificatedPair: %s", err)
    }
}
