package ethclient

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"testing"
)

var one = big.NewInt(1)

// This test is made just to ensure that a certain methods works well with go-lachesis application. It requires go-lachesis node to run
func TestCalls(t *testing.T) {
	var rawUrl = "http://127.0.0.1:18545"
	var ctx = context.Background()
	const shortEventId = "21:10:fe1423"
	client, err := Dial(rawUrl)
	require.Nil(t, err)
	chainId, err := client.ChainID(ctx)
	require.Nil(t, err)
	log.Println("chainId", chainId)

	epoch, err := client.GetCurrentEpoch(ctx)
	require.Nil(t, err)
	log.Println("epoch", epoch)

	event, err := client.GetEventHeader(ctx, shortEventId)
	require.Nil(t, err)
	log.Println("event", event)

	time, err := client.GetConsensusTime(ctx, shortEventId)
	require.Nil(t, err)
	log.Println("time", time)

	heads, err := client.GetHeads(ctx, one)
	require.Nil(t, err)
	log.Println("heads", heads)

	epochStats, err := client.GetEpochStats(ctx, one)
	require.Nil(t, err)
	log.Println("epochStats", epochStats)

	gp, err := client.GasPrice(ctx)
	require.Nil(t, err)
	log.Println("gas price ", gp)

	pv, err := client.ProtocolVersion(ctx)
	require.Nil(t, err)
	log.Println("protocol version", pv)

	snk, err := client.Syncing(ctx)
	require.Nil(t, err)
	log.Println("syncing", snk)

	bTps, err := client.GetBlockTPS(ctx, one)
	require.Nil(t, err)
	log.Println("Block TPS", bTps)

	eTps, err := client.GetEpochTPS(ctx, one)
	require.Nil(t, err)
	log.Println("Epoch TPS", eTps)

	firstBlock, err := client.BlockByNumber(ctx, one)
	require.Nil(t, err)
	trByBlockNum, err := client.GetTransactionByBlockNumberAndIndex(ctx, firstBlock.Number(), 0)
	require.Nil(t, err)
	log.Println("trByBlockNum", trByBlockNum)
	trByBlockHash, err := client.GetTransactionByBlockHashAndIndex(ctx, firstBlock.Hash(), 0)
	require.Nil(t, err)
	log.Println("trByBlockNum", trByBlockHash)
	rawTrByBlockNum, err := client.GetRawTransactionByBlockNumberAndIndex(ctx, firstBlock.Number(), 0)
	require.Nil(t, err)
	log.Println("trByBlockNum", rawTrByBlockNum)
	rawTrByBlockHash, err := client.GetRawTransactionByBlockHashAndIndex(ctx, firstBlock.Hash(), 0)
	require.Nil(t, err)
	log.Println("trByBlockNum", rawTrByBlockHash)

	//client.gas
	//client.get
}