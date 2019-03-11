package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

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
func DeployNewContract(supplier_address common.Address,
	client *ethclient.Client,
	end_time *big.Int,
	datahash [32]byte,
	account_key string,
	account_pass string) common.Address {

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(account_key), account_pass)

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := DeployEnergy(
		auth,
		client,
		supplier_address,
		end_time,
		datahash,
	)

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

	gasLimit := uint64(21000) // in units
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
