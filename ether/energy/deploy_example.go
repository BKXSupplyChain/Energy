package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

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

	// Получаем новый клиент подключенный к ноде.
	client := contracts.CreateNewClient()

	// Получаем адрес контракта, который нужно сказать поставщику.
	address := contracts.DeployNewContract(random_supplier_address,
		client,
		random_end_time,
		*random_datahash,
		key,
		"testaccountpassword")

	// Вот так можно спросить баланс контракта, да и любого другого кошелька тоже.
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// Теперь вторая важная часть, вот у нас есть контракт, но денег то на нём нет.
	// Поэтому мы туда и положим что-нибудь.

	fmt.Printf("Sleep for 2 minutes while our contract is mining.\n")
	time.Sleep(2 * time.Minute)

	value := big.NewInt(100000000000000000) // 1 eth = 1000000000000000000
	contracts.TransferEthToContract(client,
		"65FAFAF36F6B6E9A8EAC009656A5CA6FA4980F72BB0300A176B3793391905125", // Приватный ключ кошелька.
		value,
		address)
}
