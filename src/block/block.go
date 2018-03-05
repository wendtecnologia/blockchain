package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represent a block in the blockchain
type Block struct {
	timestamp         int64  //current time, when block is created
	data              []byte //valuable information
	previousBlockHash []byte //hash of the previous block
	hash              []byte //hash of the current block
}

func (block *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(block.timestamp, 10))
	headers := bytes.Join([][]byte{block.previousBlockHash, block.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	block.hash = hash[:]
}

// NewBlock creates a new block
func NewBlock(data string, previousBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), previousBlockHash, []byte{}}

	block.setHash()
	return block
}

// NewGenesisBlock create a genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}