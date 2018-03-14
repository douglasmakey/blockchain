package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewGenesisBlock create genesis block
func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

// NewBlock Create block and return
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var r bytes.Buffer
	encoder := gob.NewEncoder(&r)

	err := encoder.Encode(b)
	if err != nil {
		fmt.Printf("error: %v", err)
		return []byte{}
	}

	return r.Bytes()
}

func Deserialize(d []byte) *Block {
	var b Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Printf("error: %v", err)
		return &b
	}

	return &b
}
