package blockchain

import (
	. "github.com/wendtecnologia/blockchain/src/block"
)

// Blockchain represents the database technology, to be immutable, trackabcle
type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(data string) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]

	newBlock := NewBlock(data, previousBlock.Hash)

	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// NewBlockchain start a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
