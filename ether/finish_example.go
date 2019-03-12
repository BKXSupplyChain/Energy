package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"./contracts"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Для вызова FinishSup всё то же самое, только нужно передать чуть больше аргументов в сам вызов.

	// Получаем новый клиент подключенный к ноде.
	client := contracts.CreateNewClient()

	// Достаём адрес контракта.
	address := common.HexToAddress("0xb1299aef33c6e105e422721e4aaf4e3eb3ec4fa3")

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
}
