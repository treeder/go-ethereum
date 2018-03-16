// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package clique

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/gochain-io/gochain/common"
	"github.com/gochain-io/gochain/core"
	"github.com/gochain-io/gochain/core/types"
	"github.com/gochain-io/gochain/crypto"
	"github.com/gochain-io/gochain/ethdb"
	"github.com/gochain-io/gochain/params"
)

type testerVote struct {
	signer string
	voted  string
	auth   bool
}

// testerAccountPool is a pool to maintain currently active tester accounts,
// mapped from textual names used in the tests below to actual Ethereum private
// keys capable of signing transactions.
type testerAccountPool struct {
	accounts map[string]*ecdsa.PrivateKey
}

func newTesterAccountPool() *testerAccountPool {
	return &testerAccountPool{
		accounts: make(map[string]*ecdsa.PrivateKey),
	}
}

func (ap *testerAccountPool) sign(header *types.Header, signer string) {
	// Ensure we have a persistent key for the signer
	if ap.accounts[signer] == nil {
		ap.accounts[signer], _ = crypto.GenerateKey()
	}
	// Sign the header and embed the signature in extra data
	sig, _ := crypto.Sign(sigHash(header).Bytes(), ap.accounts[signer])
	header.Signer = sig
}

func (ap *testerAccountPool) address(account string) common.Address {
	// Ensure we have a persistent key for the account
	if ap.accounts[account] == nil {
		ap.accounts[account], _ = crypto.GenerateKey()
	}
	// Resolve and return the Ethereum address
	return crypto.PubkeyToAddress(ap.accounts[account].PublicKey)
}

// testerChainReader implements consensus.ChainReader to access the genesis
// block. All other methods and requests will panic.
type testerChainReader struct {
	db ethdb.Database
}

func (r *testerChainReader) Config() *params.ChainConfig                 { return params.AllCliqueProtocolChanges }
func (r *testerChainReader) CurrentHeader() *types.Header                { panic("not supported") }
func (r *testerChainReader) GetHeader(common.Hash, uint64) *types.Header { panic("not supported") }
func (r *testerChainReader) GetBlock(common.Hash, uint64) *types.Block   { panic("not supported") }
func (r *testerChainReader) GetHeaderByHash(common.Hash) *types.Header   { panic("not supported") }
func (r *testerChainReader) GetHeaderByNumber(number uint64) *types.Header {
	if number == 0 {
		return core.GetHeader(r.db, core.GetCanonicalHash(r.db, 0), 0)
	}
	panic("not supported")
}

// Tests that voting is evaluated correctly for various simple and complex scenarios.
func TestVoting(t *testing.T) {
	ctx := context.Background()
	// Define the various voting scenarios to test
	tests := []struct {
		epoch          uint64
		signers        []string
		voters         []string
		votes          []testerVote
		signersResults []string
		votersResults  []string
	}{
		{
			// Single signer, no votes cast
			signers:        []string{"A"},
			voters:         []string{"A"},
			votes:          []testerVote{{signer: "A"}},
			signersResults: []string{"A"},
			votersResults:  []string{"A"},
		}, {
			// Single signer, voting to add two others (only accept first)
			signers: []string{"A"},
			voters:  []string{"A"},
			votes: []testerVote{
				{signer: "A", voted: "B", auth: true},
				{signer: "B"},
				{signer: "A", voted: "C", auth: false},
				{signer: "B", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A"},
		}, {
			// Two signers, voting to add three others (only accept first two)
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: true},
				{signer: "B", voted: "C", auth: true},
				{signer: "A", voted: "D", auth: true},
				{signer: "B", voted: "D", auth: true},
				{signer: "C"},
				{signer: "A", voted: "E", auth: false},
				{signer: "B", voted: "E", auth: false},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B"},
		}, {
			// Single signer, dropping itself (weird, but one less cornercase by explicitly allowing this)
			signers: []string{"A"},
			voters:  []string{"A"},
			votes: []testerVote{
				{signer: "A", voted: "A", auth: false},
			},
			signersResults: []string{"A"},
			votersResults:  []string{},
		}, {
			// Two signers, actually needing mutual consent to drop either of them (not fulfilled)
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "B", auth: false},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A", "B"},
		}, {
			// Two signers, actually needing mutual consent to drop either of them (fulfilled)
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "B", auth: false},
				{signer: "B", voted: "B", auth: false},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A"},
		}, {
			// Three signers, two of them deciding to drop the third
			signers: []string{"A", "B", "C"},
			voters:  []string{"A", "B", "C"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B", voted: "C", auth: false},
			},
			signersResults: []string{"A", "B", "C"},
			votersResults:  []string{"A", "B"},
		}, {
			// Four signers, consensus of two not being enough to drop anyone
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B", voted: "C", auth: false},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B", "C", "D"},
		}, {
			// Four signers, consensus of three already being enough to drop someone
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "D", auth: false},
				{signer: "B", voted: "D", auth: false},
				{signer: "C", voted: "D", auth: false},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B", "C"},
		}, {
			// Authorizations are counted once per signer per target
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: true},
				{signer: "B"},
				{signer: "A", voted: "C", auth: true},
				{signer: "B"},
				{signer: "A", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A", "B"},
		}, {
			// Authorizing multiple accounts concurrently is permitted
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: true},
				{signer: "B"},
				{signer: "A", voted: "D", auth: true},
				{signer: "B"},
				{signer: "A"},
				{signer: "B", voted: "D", auth: true},
				{signer: "A"},
				{signer: "B", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B"},
		}, {
			// Deauthorizations are counted once per signer per target
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "B", auth: false},
				{signer: "B"},
				{signer: "A", voted: "B", auth: false},
				{signer: "B"},
				{signer: "A", voted: "B", auth: false},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A", "B"},
		}, {
			// Deauthorizing multiple accounts concurrently is permitted
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "A", voted: "D", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "A"},
				{signer: "B", voted: "D", auth: false},
				{signer: "C", voted: "D", auth: false},
				{signer: "A"},
				{signer: "B", voted: "C", auth: false},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B"},
		}, {
			// Votes from deauthorized signers are discarded immediately (deauth votes)
			signers: []string{"A", "B", "C"},
			voters:  []string{"A", "B", "C"},
			votes: []testerVote{
				{signer: "C", voted: "B", auth: false},
				{signer: "A", voted: "C", auth: false},
				{signer: "B", voted: "C", auth: false},
				{signer: "A", voted: "B", auth: false},
			},
			signersResults: []string{"A", "B", "C"},
			votersResults:  []string{"A", "B"},
		}, {
			// Votes from deauthorized signers are discarded immediately (auth votes)
			signers: []string{"A", "B", "C"},
			voters:  []string{"A", "B", "C"},
			votes: []testerVote{
				{signer: "C", voted: "B", auth: false},
				{signer: "A", voted: "C", auth: false},
				{signer: "B", voted: "C", auth: false},
				{signer: "A", voted: "B", auth: false},
			},
			signersResults: []string{"A", "B", "C"},
			votersResults:  []string{"A", "B"},
		}, {
			// Cascading changes are not allowed, only the account being voted on may change
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "A", voted: "D", auth: false},
				{signer: "B", voted: "C", auth: false},
				{signer: "C"},
				{signer: "A"},
				{signer: "B", voted: "D", auth: false},
				{signer: "C", voted: "D", auth: false},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B", "C"},
		}, {
			// Changes reaching consensus out of bounds (via a deauth) execute on touch
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "A", voted: "D", auth: false},
				{signer: "B", voted: "C", auth: false},
				{signer: "C"},
				{signer: "A"},
				{signer: "B", voted: "D", auth: false},
				{signer: "C", voted: "D", auth: false},
				{signer: "A"},
				{signer: "B"},
				{signer: "C", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B"},
		}, {
			// Changes reaching consensus out of bounds (via a deauth) may go out of consensus on first touch
			signers: []string{"A", "B", "C", "D"},
			voters:  []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "A", voted: "D", auth: false},
				{signer: "B", voted: "C", auth: false},
				{signer: "C"},
				{signer: "A"},
				{signer: "B", voted: "D", auth: false},
				{signer: "C", voted: "D", auth: false},
				{signer: "A"},
				{signer: "B", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B", "C", "D"},
			votersResults:  []string{"A", "B", "C"},
		}, {
			// Ensure that pending votes don't survive authorization status changes. This
			// corner case can only appear if a signer is quickly added, removed and then
			// readded (or the inverse), while one of the original voters dropped. If a
			// past vote is left cached in the system somewhere, this will interfere with
			// the final signer outcome.
			signers: []string{"A", "B", "C", "D", "E"},
			voters:  []string{"A", "B", "C", "D", "E"},
			votes: []testerVote{
				{signer: "A", voted: "F", auth: true}, // Authorize F, 3 votes needed
				{signer: "B", voted: "F", auth: true},
				{signer: "C", voted: "F", auth: true},
				{signer: "D", voted: "F", auth: false}, // Deauthorize F, 3 votes needed (leave A's previous vote "unchanged")
				{signer: "E", voted: "F", auth: false},
				{signer: "B"},
				{signer: "C"},
				{signer: "D", voted: "F", auth: true}, // Almost authorize F as a voter, 2/3 votes needed
				{signer: "E", voted: "F", auth: true},
				{signer: "B", voted: "A", auth: false}, // Deauthorize A as a voter, 3 votes needed
				{signer: "C", voted: "A", auth: false},
				{signer: "D", voted: "A", auth: false},
				{signer: "E"},
				{signer: "B", voted: "F", auth: true}, // Finish authorizing F as a voter, 3/3 votes needed
			},
			//results: []string{"B", "C", "D", "E", "F"},
			signersResults: []string{"A", "B", "C", "D", "E", "F"},
			votersResults:  []string{"B", "C", "D", "E", "F"},
		}, {
			// Epoch transitions reset all votes to allow chain checkpointing
			epoch:   3,
			signers: []string{"A", "B"},
			voters:  []string{"A", "B"},
			votes: []testerVote{
				{signer: "A", voted: "C", auth: true},
				{signer: "B"},
				{signer: "A"}, // Checkpoint block, (don't vote here, it's validated outside of snapshots)
				{signer: "B", voted: "C", auth: true},
			},
			signersResults: []string{"A", "B"},
			votersResults:  []string{"A", "B"},
		},
	}
	// Run through the scenarios and test them
	for i, tt := range tests {
		// Create the account pool and generate the initial set of signers
		accounts := newTesterAccountPool()

		signers := make([]common.Address, len(tt.signers))
		voters := make([]common.Address, len(tt.voters))
		for j, signer := range tt.signers {
			signers[j] = accounts.address(signer)
		}
		for j, voter := range tt.voters {
			voters[j] = accounts.address(voter)
		}
		for j := 0; j < len(signers); j++ {
			for k := j + 1; k < len(signers); k++ {
				if bytes.Compare(signers[j][:], signers[k][:]) > 0 {
					signers[j], signers[k] = signers[k], signers[j]
				}
			}
		}
		for j := 0; j < len(voters); j++ {
			for k := j + 1; k < len(voters); k++ {
				if bytes.Compare(voters[j][:], voters[k][:]) > 0 {
					voters[j], voters[k] = voters[k], voters[j]
				}
			}
		}
		// Create the genesis block with the initial set of signers
		genesis := &core.Genesis{
			ExtraData: make([]byte, extraVanity),
			Signers:   signers,
			Voters:    voters,
			Signer:    make([]byte, extraSeal),
		}
		// Create a pristine blockchain with the genesis injected
		db, _ := ethdb.NewMemDatabase()
		genesis.Commit(db)

		// Assemble a chain of headers from the cast votes
		headers := make([]*types.Header, len(tt.votes))
		for j, vote := range tt.votes {
			headers[j] = &types.Header{
				Number:   big.NewInt(int64(j) + 1),
				Time:     big.NewInt(int64(j) * int64(blockPeriod)),
				Coinbase: accounts.address(vote.voted),
				Signer:   make([]byte, extraSeal),
			}
			if j > 0 {
				headers[j].ParentHash = headers[j-1].Hash()
			}
			if vote.auth {
				copy(headers[j].Nonce[:], nonceAuthVote)
			}
			accounts.sign(headers[j], vote.signer)
		}
		// Pass all the headers through clique and ensure tallying succeeds
		head := headers[len(headers)-1]

		snap, err := New(&params.CliqueConfig{Epoch: tt.epoch}, db).snapshot(ctx, &testerChainReader{db: db}, head.Number.Uint64(), head.Hash(), headers)
		if err != nil {
			t.Errorf("test %d: failed to create voting snapshot: %v", i, err)
			continue
		}
		// Verify the final list of signers against the expected ones
		signers = make([]common.Address, len(tt.signersResults))
		for j, signer := range tt.signersResults {
			signers[j] = accounts.address(signer)
		}
		for j := 0; j < len(signers); j++ {
			for k := j + 1; k < len(signers); k++ {
				if bytes.Compare(signers[j][:], signers[k][:]) > 0 {
					signers[j], signers[k] = signers[k], signers[j]
				}
			}
		}
		signersResult := snap.signers()
		if len(signersResult) != len(signers) {
			t.Errorf("test %d: signers mismatch: have %x, want %x", i, signersResult, signers)
			continue
		}
		for j := 0; j < len(signersResult); j++ {
			if !bytes.Equal(signersResult[j][:], signers[j][:]) {
				t.Errorf("test %d, signer %d: signer mismatch: have %x, want %x", i, j, signersResult[j], signers[j])
			}
		}
		// Verify the final list of voters against the expected ones
		voters = make([]common.Address, len(tt.votersResults))
		for j, voter := range tt.votersResults {
			voters[j] = accounts.address(voter)
		}
		for j := 0; j < len(voters); j++ {
			for k := j + 1; k < len(voters); k++ {
				if bytes.Compare(voters[j][:], voters[k][:]) > 0 {
					voters[j], voters[k] = voters[k], voters[j]
				}
			}
		}
		votersResult := snap.voters()
		if len(votersResult) != len(voters) {
			t.Errorf("test %d: voters mismatch: have %x, want %x", i, votersResult, voters)
			continue
		}
		for j := 0; j < len(votersResult); j++ {
			if !bytes.Equal(votersResult[j][:], voters[j][:]) {
				t.Errorf("test %d, voter %d: voter mismatch: have %x, want %x", i, j, votersResult[j], voters[j])
			}
		}
	}
}
