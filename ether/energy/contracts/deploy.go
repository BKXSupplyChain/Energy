package contracts

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Returns an address of the new contract.
func DeployNewContract(supplier_address common.Address,
	end_time *big.Int,
	datahash [32]byte,
	eth_node_address string,
	account_key string,
	account_pass string) common.Address {

	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial(eth_node_address)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(account_key), account_pass)

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := DeployEnergy(
		auth,
		blockchain,
		supplier_address,
		end_time,
		datahash,
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)

	return address
}
