package send
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/rpc"
)


func sendProposal(proposal types.Proposal) {
    client, err := rpc.DialHTTP("tcp", ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("Proposal.OnProposalReceived", proposal, {})
    if (err != nil) {
        log.Fatalf("Error in Proposal: %s", err)
    }
}
