package main

import (
	"math/big"

	"./contracts"
	"github.com/ethereum/go-ethereum/common"
)

// Грубо говоря кошелёк покупателя
const key = `{"address":"89d277c7ff0b4ed696db62da81cf2b910a32a258",
			  "crypto":{"cipher":"aes-128-ctr",
						  "ciphertext":"f44cc47a068bb9fdeef2413690bfa15379d7489d543f8b65f658fc17dcae90f0",
						  "cipherparams":{"iv":"7abf24c6e1e2bb449db4500142736991"},
						  "kdf":"scrypt",
						  "kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"9e1545d62a8552b16518ba747689e4f53d4bcb9dc6008deca375a672bcd1f5bf"},
						  "mac":"a08d3f3ebe98510d7bc6f4c0093d22031799d19a2d48afecc123b9ab43a3544e"},
						  "id":"9642a72a-b08b-408a-9d93-6903e45f1858",
						  "version":3}`

func main() {
	// Здесь нужно взять адрес поставщика
	random_supplier_address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	// Время окончания
	random_end_time := new(big.Int)
	// Хэш: посмотри как он используется в самом контракте при разрыве или спроси у Артемия как его насчитать
	random_datahash := new([32]byte)

	// Получаем адрес контракта, который нужно сказать поставщику
	address := contracts.DeployNewContract(random_supplier_address,
		random_end_time,
		*random_datahash,
		"https://rinkeby.infura.io/v3/c5035536f7a0467299308c86b3eb9dbc",
		key,
		"testaccountpassword")

	// Чтобы go не ругался
	address = address

	// TODO: нужно положить по адресу контракта сумму, о которой договорились.
}
