package model

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vladmarchuk90/go-api-server-test/contract"
)

type Group struct {
	Name    string     `json:"name"`
	Indexes []*big.Int `json:"indexes"`
}

type Index struct {
	Name              string   `json:"name"`
	EthPriceInWei     *big.Int `json:"ethpriceinwei"`
	UsdPriceInCents   *big.Int `json:"usdpriceincents"`
	UsdCapitalization *big.Int `json:"usdcapitalization"`
	PercentageChange  *big.Int `json:"percentagechange"`
}

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
func GetGroups() ([]*big.Int, error) {

	groupIds, err := smartContract.GetGroupIds(nil)
	if err != nil {
		return groupIds, fmt.Errorf("failed to get groups: %v", err)
	}

	return groupIds, nil
}

func GetGroup(groupIdString string) (*Group, error) {

	var group *Group

	groupId := new(big.Int)
	groupId, ok := groupId.SetString(groupIdString, 10)
	if !ok {
		return group, fmt.Errorf("failed to convert group Id parameter to big Int: %v", groupIdString)
	}

	contractGroup, err := smartContract.GetGroup(nil, groupId)
	if err != nil {
		return group, fmt.Errorf("failed to get group by id: %v, error: %v", groupId, err)
	}

	_group := Group(contractGroup)

	group = &_group
	return group, nil
}

func GetIndex(indexIdString string) (*Index, error) {

	var index *Index

	indexId := new(big.Int)
	indexId, ok := indexId.SetString(indexIdString, 10)
	if !ok {
		return index, fmt.Errorf("failed to convert group Id parameter to big Int: %v", indexIdString)
	}

	contractIndex, err := smartContract.GetIndex(nil, indexId)
	if err != nil {
		return index, fmt.Errorf("failed to get index by id: %v, error: %v", indexId, err)
	}

	_index := Index(contractIndex)
	index = &_index

	return index, nil
}

func GetBlock(blockParameter string) (*types.Header, error) {

	var blockInfo *types.Header

	if blockParameter == "latest" {
		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			return blockInfo, fmt.Errorf("failed to get latest block: %v,", err)
		}

		blockInfo = block.Header()
		return blockInfo, nil
	} else {
		// Trying to understand what kind of parameter hash or number
		// Firstly check if it's number
		blockNumber := new(big.Int)
		blockNumber, ok := blockNumber.SetString(blockParameter, 10)

		if ok {

			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				return blockInfo, fmt.Errorf("failed to get block by number: %v, error: %v", blockNumber, err)
			}
			blockInfo = block.Header()
			return blockInfo, nil

		} else {

			// Maybe hash
			hash := common.HexToHash(blockParameter)
			block, err := client.BlockByHash(context.Background(), hash)
			if err != nil {
				return blockInfo, fmt.Errorf("failed to get block by hash: %v, error: %v", hash, err)
			}
			blockInfo = block.Header()
			return blockInfo, nil
		}
	}
}
