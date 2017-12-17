package main

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

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.previousBlockHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.hash = hash[:]
}

func new(data string, previousBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), previousBlockHash, []byte{}}
	block.setHash()
	return block
}
