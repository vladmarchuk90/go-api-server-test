package model

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vladmarchuk90/go-api-server-test/contract"
)

var client *ethclient.Client
var smartContract *contract.SmartcontractCaller

func ConnectToEthereumSmartContract() {

	var err error

	// Firstly get Ethereum contract bindings, so we can call methods of smart contract on  Ropsten test network
	client, err = ethclient.Dial("https://ropsten.infura.io/v3/6fb9b00bf7a045b280ea86355f9deeba")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	address := common.HexToAddress("0x4f7f1380239450AAD5af611DB3c3c1bb51049c29")
	smartContract, err = contract.NewSmartcontractCaller(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a smart contract: %v", err)
	}
}

func CloseConnection() {
	client.Close()
}

// Respective methods to api/comtroller methods
func GetGroups() (groupIds []*big.Int) {

	groupIds, err := smartContract.GetGroupIds(nil)
	if err != nil {
		log.Printf("Failed to get groups: %v", err)
		return
	}

	return
}

func GetGroup(groupIdString string) (group struct {
	Name    string
	Indexes []*big.Int
}) {

	groupId := new(big.Int)
	groupId, ok := groupId.SetString(groupIdString, 10)
	if !ok {
		log.Printf("Failed to convert group Id parameter in big Int: %v", groupIdString)
		return
	}

	group, err := smartContract.GetGroup(nil, groupId)
	if err != nil {
		log.Printf("Failed to get group by id: %v, error: %v", groupId, err)
		return
	}

	return
}

func GetIndex(indexIdString string) (index struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}) {

	indexId := new(big.Int)
	indexId, ok := indexId.SetString(indexIdString, 10)
	if !ok {
		log.Printf("Failed to convert index Id parameter in big Int: %v", indexIdString)
		return
	}

	index, err := smartContract.GetIndex(nil, indexId)
	if err != nil {
		log.Printf("Failed to get index by id: %v, error: %v", indexId, err)
		return
	}

	return
}

func GetBlock(blockParameter string) (block *types.Header) {

	if blockParameter == "latest" {
		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Printf("Failed to get latest block: %v,", err)
		}
		return block.Header()
	} else {
		// Trying to understand what kind of parameter hash or number
		// Firstly check if it's number
		blockNumber := new(big.Int)
		blockNumber, ok := blockNumber.SetString(blockParameter, 10)

		if ok {

			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				log.Printf("Failed to get block by number: %v, error: %v", blockNumber, err)
			}
			return block.Header()

		} else {

			// Maybe hash
			hash := common.HexToHash(blockParameter)
			block, err := client.BlockByHash(context.Background(), hash)
			if err != nil {
				log.Printf("Failed to get block by hash: %v, error: %v", hash, err)
			}
			return block.Header()

		}
	}
}
