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

func correctDivision (a big.Int, b big.Int) (*big.Float){
	num_float1 := new(big.Float).SetInt(&a)
	num_float2 := new(big.Float).SetInt(&b)
	return new(big.Float).Quo(num_float1, num_float2)
}

func compareFloats (num1 big.Float, num2 float64) int {
	numBigFloat2 := new(big.Float).SetFloat64(num2)
	return num1.Cmp(numBigFloat2)
}

func provider(socketID string, userID string) {

	//TODO get from provider (r, s, amountGet, contractID = socket.ActiveContract)
	var contractID string // = socket.ActiveContract

	var contract types.Contract
	df.Get(&contract, contractID)

	var socket types.SocketInfo
	db.Get(&socket, contract.Socket)

	var curentProposal types.Proposal
	db.Get(&curentProposal, socket.ActiveProposal)

	rawAllEnergy := db.getEnergy(socketID) //uint64
	allEnergy := new(big.Int).SetUint64(rawAllEnergy)

	var amountWant = new(big.Int)
	amountWant.Mul(allEnergy, curentProposal.Price)

	var msg string = amountWant.String() + socket.ActiveContract

	hash := sha256.Sum256([]byte(msg))
	valid := ecdsa.Verify(HexToECDSA(socket.NeighborKey), hash[:], r, s)

	var absError big.Int = curentProposal.AbsError
	var relError float64 = curentProposal.RelError

	var delta big.Int
	delta.Sub(amountWant, amountGet)
	var absDelta big.Int
	absDelta.Abs(&delta)

	relerr := correctDivision(*amountGet, *amountWant)


	if (valid) {
		if (delta.Cmp(&absError) <= 0 && compareFloats(*relerr, relError) >= 0) {

			var signature string = r.String() + " " + s.String()

			contract.LastSertificate.Signature = signature
			contract.LastSertificate.Amount = amountGet.bytes()

		} else {

			socket.ActiveProposal = ""

			client := contracts.CreateNewClient()

			// Достаём адрес контракта.
			address := common.HexToAddress(contract.EthAddress)

			// Приватный ключ для подписи вызова.
			rawPrivateKey := "65FAFAF36F6B6E9A8EAC009656A5CA6FA4980F72BB0300A176B3793391905125" //TODO WHETE GET IT?

			// Получаем последний сертификат
			var err bool
			strPair := strings.Split(contract.LastSertificate.Signature, " ")
			r, err = r.SetString(strPair[0], 10)
			if !err {
				//error
			}
			s, err = s.SetString(strPair[1], 10)
			if !err {
				//error
			}


			var lastAmount = new(big.Int).SetBytes(contract.LastSertificate.Amount)
			contracts.FinishSup(client, address, rawPrivateKey, lastAmount , HexToECDSA(socket.NeighborKey),  r, s)

			fmt.Printf("Waiting for a minute while our finish call is deploying.\n")
			time.Sleep(time.Minute)

			// Проверяем, что денег на контракте не осталось.
			balance, err := client.BalanceAt(context.Background(), address, nil)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(balance)
		}
	}

}

