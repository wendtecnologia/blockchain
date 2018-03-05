package blockchain

import (
	. "github.com/wendtecnologia/blockchain/src/block"
)

// Blockchain represents the database technology, to be immutable, trackabcle
type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) addBlock(data string) {
	previousBlock := bc.blocks[len(bc.blocks)-1]

	newBlock := NewBlock(data, previousBlock.hash)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain start a new bc
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
