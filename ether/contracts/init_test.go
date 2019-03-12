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

//Test initial contract gets set up correctly
func TestGetData(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 7900000)

	random_supplier_address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	random_end_time := new(big.Int)
	random_datahash := new([32]byte)

	//Deploy contract
	_, _, contract, _ := DeployEnergy(
		auth,
		blockchain,
		random_supplier_address,
		random_end_time,
		*random_datahash,
	)

	// commit all pending transactions
	blockchain.Commit()

	if got, _ := contract.Consumer(nil); got != auth.From {
		t.Errorf("Got wrong consumer.")
	}
	if got, _ := contract.Supplier(nil); got != random_supplier_address {
		t.Errorf("Got wrong supplier.")
	}
	if got, _ := contract.EndTime(nil); got.Cmp(random_end_time) != 0 {
		t.Errorf("Got wrong endTime.")
	}
	if got, _ := contract.DataHash(nil); got != *random_datahash {
		t.Errorf("Got wrong dataHash.")
	}
}
