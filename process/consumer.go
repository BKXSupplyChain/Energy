import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
	"github.com/BKXSupplyChain/Energy/db"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/BKXSupplyChain/Energy/tree/master/ether/contracts"
)




func consumer(socketID string, userID string) {

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

	//TODO send to provider (r, s, amount, socket.ActiveContract)
}
