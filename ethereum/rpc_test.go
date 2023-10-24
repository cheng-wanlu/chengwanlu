package eth_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rpc"
)

func TestRPC(t *testing.T) {

	rawurl := "https://mainnet.infura.io/v3/5c17ecf14e0d4756aa81b6a1154dc599"

	// rawClient, err := rpc.DialHTTP(rawurl)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// ethClient := ethclient.NewClient(rawClient)

	ethClient, err := ethclient.Dial(rawurl)
	if err != nil {
		t.Fatal(err)
	}
	defer ethClient.Close()

	// APIs: https://docs.infura.io/networks/ethereum/json-rpc-methods

	blocknumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	blockNumber := big.NewInt(int64(blocknumber))
	t.Logf("当前区块高度: %d\n", blockNumber)

	// 币安ETH地址
	binanceAddress, err := hex.DecodeString("BE0eB53F46cd790Cd13851d5EFf43D12404d33E8")
	if err != nil {
		t.Fatal(err)
	}
	balance, err := ethClient.BalanceAt(context.Background(), common.Address(binanceAddress), blockNumber)
	if err != nil {
		t.Fatal(err)
	}
	balancef, _ := balance.Float64()
	t.Logf("查询币安 ETH 地址余额: %f ETH\n", balancef/1e18)

	block, err := ethClient.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("查询区块转账记录, 区块哈希: %s\n", block.Hash())
	for _, tx := range block.Transactions() {
		from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			t.Fatal(err)
		}

		value, _ := tx.Value().Float64()
		t.Logf("交易哈希: %s \nFROM: %s TO:%s %f ETH\n", tx.Hash(), from, tx.To(), value/1e18)
	}

	// 查询 ERC20 USDT余额
	usdtContractAddress := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")

	contract, err := NewToken(usdtContractAddress, ethClient)
	if err != nil {
		t.Fatal(err)
	}

	binanceUSDTAddress := common.HexToAddress("0xDFd5293D8e347dFe59E90eFd55b2956a1343963d")
	usdtBalance, err := contract.BalanceOf(&bind.CallOpts{}, binanceUSDTAddress)
	if err != nil {
		t.Fatal(err)
	}
	usdtBalancef, _ := usdtBalance.Float64()
	t.Logf("查询币安 ERC20, USDT 地址余额: %f USDT\n", usdtBalancef/1e6)

	// 查询 ERC20 USDT转账
	query := ethereum.FilterQuery{
		Addresses: []common.Address{usdtContractAddress},
		Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
	}

	// Filter the logs
	transferLogs, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("查询 ERC20 USDT 转账记录, 区块哈希: %s\n", block.Hash())
	for _, transferLog := range transferLogs {
		from := common.HexToAddress(transferLog.Topics[1].Hex())
		to := common.HexToAddress(transferLog.Topics[2].Hex())
		amount := new(big.Int)
		amount.SetBytes(transferLog.Data)
		valuef, _ := amount.Float64()
		t.Logf("交易哈希: %s \nFROM: %s TO:%s %f USDT\n", transferLog.TxHash, from, to, valuef/1e6)
	}
}
