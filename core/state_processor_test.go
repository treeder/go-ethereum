package core

import (
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/gochain-io/gochain/common"
	"github.com/gochain-io/gochain/core/state"
	"github.com/gochain-io/gochain/core/types"
	"github.com/gochain-io/gochain/core/vm"
	"github.com/gochain-io/gochain/crypto"
	"github.com/gochain-io/gochain/params"
)

func BenchmarkStateProcessor_Process(b *testing.B) {
	key, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(key.PublicKey)
	funds := big.NewInt(1000000000)

	genesis := &Genesis{
		Config:     params.TestChainConfig,
		Difficulty: big.NewInt(1),
		Alloc:      GenesisAlloc{address: {Balance: funds}},
	}
	signer := types.NewEIP155Signer(genesis.Config.ChainId)

	bc := newTestBlockChainWithGenesis(true, true, genesis)
	defer bc.Stop()
	cfg := vm.Config{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		txs := make([]*types.Transaction, 1000)
		for i := 0; i < 1000; i++ {
			tx := types.NewTransaction(uint64(i), common.Address{}, big.NewInt(100), 100000, big.NewInt(1), nil)
			tx, _ = types.SignTx(tx, signer, key)
			txs[i] = tx
		}

		block := types.NewBlock(&types.Header{
			GasLimit: bc.GasLimit(),
		}, txs, nil, nil)
		statedb, err := state.New(bc.CurrentBlock().Root(), bc.stateCache)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		_, _, _, err = bc.Processor().Process(block, statedb, cfg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestStateProcessor(b *testing.T) {
	start := time.Now()
	key, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(key.PublicKey)
	funds := big.NewInt(1000000000)

	genesis := &Genesis{
		Config:     params.TestChainConfig,
		Difficulty: big.NewInt(1),
		Alloc:      GenesisAlloc{address: {Balance: funds}},
	}
	signer := types.NewEIP155Signer(genesis.Config.ChainId)

	bc := newTestBlockChainWithGenesis(true, true, genesis)
	log.Printf("newTestBlockchain duration: %s", time.Since(start))
	defer bc.Stop()
	cfg := vm.Config{}
	// statedb, err := state.New(bc.CurrentBlock().Root(), bc.stateCache)
	// 	if err != nil {
	// 		b.Fatal(err)
	// 	}
	// for i := 0; i < b.N; i++ {
	txs := []*types.Transaction{}
	start = time.Now()
	var newTxDur time.Duration
	var signDur time.Duration
	for i := 0; i < 1000; i++ {
		start2 := time.Now()
		tx := types.NewTransaction(uint64(i), common.Address{}, big.NewInt(100), 100000, big.NewInt(1), nil)
		newTxDur = newTxDur + time.Since(start2)
		start2 = time.Now()
		tx, _ = types.SignTx(tx, signer, key)
		signDur = signDur + time.Since(start2)
		// txs = append(txs, types.NewTransaction(uint64(i), common.Address{}, big.NewInt(1), uint64(21000), big.NewInt(21000), nil))
		txs = append(txs, tx)
	}
	log.Printf("tx creation duration: %s newTxDur: %s signTxDur: %s", time.Since(start), newTxDur, signDur)
	block := types.NewBlock(&types.Header{
		GasLimit: bc.GasLimit(),
	}, txs, nil, nil)
	statedb, err := state.New(bc.CurrentBlock().Root(), bc.stateCache)
	if err != nil {
		b.Fatal(err)
	}
	start = time.Now()
	_, _, _, err = bc.Processor().Process(block, statedb, cfg)
	if err != nil {
		b.Fatal(err)
	}
	log.Printf("process() duration: %s", time.Since(start))
	// }
}
