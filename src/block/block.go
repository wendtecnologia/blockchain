package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represent a block in the blockchain
type Block struct {
	Timestamp         int64  //current time, when block is created
	Data              []byte //valuable information
	PreviousBlockHash []byte //hash of the previous block
	Hash              []byte //hash of the current block
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreviousBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	block.Hash = hash[:]
}

// NewBlock creates a new block
func NewBlock(data string, previousBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), previousBlockHash, []byte{}}

	block.SetHash()
	return block
}

// NewGenesisBlock create a genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
