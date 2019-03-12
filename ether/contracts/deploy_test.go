package contracts

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// Test energy contract gets deployed correctly
func TestDeployEnergy(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 7900000) // Just a random number

	random_supplier_address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	random_end_time := new(big.Int)
	random_datahash := new([32]byte)

	//Deploy contract
	address, _, _, err := DeployEnergy(
		auth,
		blockchain,
		random_supplier_address,
		random_end_time,
		*random_datahash,
	)

	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the Energy contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

	// Weak point: check if all pointers are cleared
}
