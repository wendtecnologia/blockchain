package blockchain

import (
	. "github.com/wendtecnologia/blockchain/src/block"
)

// Blockchain represents the database technology, to be immutable, trackabcle
type Blockchain struct {
	blocks []*Block
}

func (blockchain *Blockchain) addBlock(data string) {
	previousBlock := blockchain.blocks[len(blockchain.blocks)-1]

	newBlock := NewBlock(data, previousBlock.hash)

	blockchain.blocks = append(blockchain.blocks, newBlock)
}

// NewBlockchain start a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
