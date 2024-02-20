package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data"
	"log"
	"math/big"
	"strings"
)

type EventListener struct {
	Transfer        data.ITransfer
	ClientUrl       string
	ContractAddress string
	ABI             string
}

func NewEventListener(t data.ITransfer, url, address, abi string) *EventListener {
	return &EventListener{
		Transfer:        t,
		ClientUrl:       url,
		ContractAddress: address,
		ABI:             abi,
	}
}

func (e *EventListener) Run() {
	client, err := ethclient.Dial(e.ClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(e.ContractAddress)},
	}
	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(e.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLog := <-ch:
			var Event struct {
				Value *big.Int
			}
			err = contractABI.UnpackIntoInterface(&Event, "Transfer", eventLog.Data)
			if err != nil {
				log.Println("Failed to unpack")
				continue
			}
			transfer := data.Transfer{
				FromAddress: common.BytesToAddress(eventLog.Topics[1].Bytes()).Hex(),
				ToAddress:   common.BytesToAddress(eventLog.Topics[2].Bytes()).Hex(),
				Value:       Event.Value.String(),
			}
			err = e.Transfer.Add(transfer)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
