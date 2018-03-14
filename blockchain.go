package main

import (
	"fmt"
)

type Blockchain struct {
	pointer []byte
	db      DB
}

func NewBlockchain(dbPathOrName string) *Blockchain {
	// Instance Persistence
	persistence := new(Bolt)
	pointer, err := persistence.Init(dbPathOrName)
	if err != nil {
		fmt.Printf("error: %v ", err)
	}

	return &Blockchain{pointer, persistence}
}

func (bc *Blockchain) AddBlock(data string) {
	lastHash, err := bc.db.GetLast()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	newBlock := NewBlock(data, lastHash)
	newPointer, err := bc.db.Update(newBlock)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	bc.pointer = newPointer
}

type BlockchainManager struct {
	bc             *Blockchain
	currentPointer []byte
}

func (bc *Blockchain) Manager() *BlockchainManager {
	bm := &BlockchainManager{bc, bc.pointer}

	return bm
}

func (bm *BlockchainManager) Current() *Block {
	var block *Block

	block, err := bm.bc.db.Get(bm.currentPointer)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return block
}

func (bm *BlockchainManager) Next() *Block {
	var block *Block

	block, err := bm.bc.db.Get(bm.currentPointer)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	bm.currentPointer = block.PrevBlockHash

	return block
}
