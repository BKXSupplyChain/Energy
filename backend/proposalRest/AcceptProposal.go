package proposalRest

import (
    "context"
	"fmt"
	"log"
	"time"

	"github.com/BKXSupplyChain/Energy/ether/contracts"
	"github.com/ethereum/go-ethereum/common"
)


func AcceptProposal(hexAddrConsemer string, hexAddrSupplier string, endTime int32, value int32, datahash string) {
    // Завершаем старый контракт:
	// Для вызова FinishSup всё то же самое, только нужно передать чуть больше аргументов в сам вызов.

	// Получаем новый клиент подключенный к ноде.
	client := contracts.CreateNewClient()

	// Достаём адрес контракта.
	address := common.HexToAddress(hexAddrConsumer)

	// Приватный ключ для подписи вызова.
	rawPrivateKey := "65FAFAF36F6B6E9A8EAC009656A5CA6FA4980F72BB0300A176B3793391905125"

	contracts.FinishExp(client, address, rawPrivateKey)

	fmt.Printf("Waiting for a minute while our finish call is deploying.\n")
	time.Sleep(time.Minute)

	// Проверяем, что денег на контракте не осталось.
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

    // Заключаем новый контракт
	// Всё что нужно вытащить из базы данных и/или договора о поставке:
	// Здесь нужно взять адрес поставщика
	supplier_address := common.HexToAddress(hexAddrSupplier)
	// Время окончания
	// Приватный ключ от кошелька потребителя.
	rawPrivateKey := "65FAFAF36F6B6E9A8EAC009656A5CA6FA4980F72BB0300A176B3793391905125"

	// Получаем новый клиент подключенный к ноде.
	client := contracts.CreateNewClient()

	// Получаем адрес контракта, который нужно сказать поставщику.
	address := contracts.DeployNewContract(supplier_address,
		client,
		endTime,
		*datahash,
		value,
		rawPrivateKey)

	fmt.Printf("Waiting for a minute while our contract is deploying.\n")
	time.Sleep(time.Minute)
}

