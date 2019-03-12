package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const eth_node_address = "https://rinkeby.infura.io/v3/c5035536f7a0467299308c86b3eb9dbc"

// Connects to the ethereum node and returns the running client.
func CreateNewClient() *ethclient.Client {
	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial(eth_node_address)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	return client
}

// Returns an address of the new contract.
func DeployNewContract(supplierAddress common.Address,
	client *ethclient.Client,
	endTime *big.Int,
	datahash [32]byte,
	value *big.Int,
	rawPrivateKey string) common.Address {

	privateKey, err := crypto.HexToECDSA(rawPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := DeployEnergy(
		auth,
		client,
		supplierAddress,
		endTime,
		datahash,
	)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(tx.Hash().Hex()) // Deployment transaction.
	_ = tx

	fmt.Printf("Contract pending deploy: 0x%x\n", address)

	return address
}

// Transfers money into the contract.
func TransferEthToContract(client *ethclient.Client,
	raw_private_key string,
	value *big.Int,
	toAddress common.Address) {

	privateKey, err := crypto.HexToECDSA(raw_private_key)
	if err != nil {
		log.Fatal(err)
	}

	// Getting public key from private. Magic.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Just gettin nonce.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
