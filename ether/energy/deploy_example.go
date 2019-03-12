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

func main() {
	// Всё что нужно вытащить из базы данных и/или договора о поставке:
	// Здесь нужно взять адрес поставщика
	random_supplier_address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	// Время окончания
	randomEndTime := new(big.Int)
	// Хэш: посмотри как он используется в самом контракте при разрыве или спроси у Артемия как его насчитать
	randomDatahash := new([32]byte)
	// Сумма, которую хотим положить в контракт.
	value := big.NewInt(123456789) // 1 eth = 1000000000000000000
	// Приватный ключ от кошелька потребителя.
	rawPrivateKey := "65FAFAF36F6B6E9A8EAC009656A5CA6FA4980F72BB0300A176B3793391905125"

	// Получаем новый клиент подключенный к ноде.
	client := contracts.CreateNewClient()

	// Получаем адрес контракта, который нужно сказать поставщику.
	address := contracts.DeployNewContract(random_supplier_address,
		client,
		randomEndTime,
		*randomDatahash,
		value,
		rawPrivateKey)

	fmt.Printf("Waiting for a minute while our contract is deploying.\n")
	time.Sleep(time.Minute)

	// Вот так можно спросить баланс контракта, да и любого другого кошелька тоже.
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}
