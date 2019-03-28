package client
  
import (
    "encoding/binary"
    "errors"
    "github.com/BKXSupplyChain/Energy/db"
    "github.com/BKXSupplyChain/Energy/types"
    "net/rpc"
)

func sendProposal(proposalMsg types.ProposalMessage, userPublicKey string, socket types.SocketInfo) {
    client, err := rpc.DialHTTP("tcp", socket.NeighborAddr + ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
    err = client.Call("types.ProposalMsg.OnProposalReceived", proposalMsg, {})
    if (err != nil) {
        log.Fatalf("Error in ProposalMessage: %s", err)
    }
}


func server(userID string) {
	var user types.UserData
	db.Get(&user, userID)

	for ; true; {
		for i := 0; i < len(user.Sockets); i++ {
			consumer(user.Sockets[i], userID)
		}
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)

	}
}

func consumer(socketID string, userID string) types.Certificate {
  
	var socket types.SocketInfo
	db.Get(&socket, socketID)

	var curentProposal types.Proposal
	db.Get(&curentProposal, socket.ActiveProposal)

	rawAllEnergy := db.getEnergy(socketID) //uint64
	allEnergy := new(big.Int).SetUint64(rawAllEnergy)

	var amount = new(big.Int)
	amount.Mul(allEnergy, curentProposal.Price)

	var msg string = amount.String() + socket.ActiveContract

	hash := sha256.Sum256([]byte(msg))

	var user types.UserData
	db.Get(&user, userID)

	r, s, err := ecdsa.Sign(rand.Reader, HexToECDSA(user.PrivateKey), hash[:])

	if err != nil {
		  panic(err)
	}
  
  var res types.Certificate
  
  res.r = r.Bytes()
  res.s = s.Bytes()
  res.Amount = amount.Bytes()
  res.ActiveContract = socket.ActiveContract
  return res
}

func sendCertificate(userID string) {
    var user types.UserData
    db.Get(&user, UserID)
  
    client, err := rpc.DialHTTP("tcp", /*TODO address*/ ":30")
    if (err != nil) {
        log.Fatalf("Error in dialing: %s", err)
    }
  
    for ; true; {
		    for i := 0; i < len(user.Sockets); i++ {
           certificate := consumer(user.Sockets[i], userID)
		    }
        err = client.Call("types.Certificate."/*TODO name of function of checking if all is correct*/, certif, {})
        if (err != nil) {
           log.Fatalf("Error in CertificatedPair: %s", err)
        }
        time.Sleep(time.Second)
		    time.Sleep(time.Second)
		    time.Sleep(time.Second)
		    time.Sleep(time.Second)
        time.Sleep(time.Second)
    }
}
