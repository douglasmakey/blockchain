package main

import (
	"github.com/boltdb/bolt"
)

// DB interface for implementation of database
type DB interface {
	Init(path string) ([]byte, error)
	Get(hash []byte) (*Block, error)
	GetLast() ([]byte, error)
	Update(block *Block) ([]byte, error)
	Close()
}

// Bolt implementation db with Bolt
type Bolt struct {
	DB *bolt.DB
}

func (b *Bolt) Close() {
	b.DB.Close()
}

func (b *Bolt) Init(dbPathOrName string) ([]byte, error) {
	var pointer []byte

	db, err := bolt.Open(dbPathOrName, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte("blocks"))
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("last"), genesis.Hash)
			if err != nil {
				return err
			}
			pointer = genesis.Hash
		} else {
			pointer = b.Get([]byte("last"))
		}

		return nil
	})

	b.DB = db

	return pointer, nil
}

func (b *Bolt) Update(block *Block) ([]byte, error) {
	var pointer []byte

	err := b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		err := b.Put(block.Hash, block.Serialize())
		err = b.Put([]byte("last"), block.Hash)
		if err != nil {
			return err
		}
		pointer = block.Hash

		return nil
	})

	if err != nil {
		return nil, err
	}

	return pointer, nil
}

func (b *Bolt) Get(hash []byte) (*Block, error) {
	var block *Block

	err := b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		encodedBlock := b.Get(hash)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return block, nil

}

func (b *Bolt) GetLast() ([]byte, error) {
	var lastHash []byte

	err := b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		lastHash = b.Get([]byte("last"))

		return nil
	})

	if err != nil {
		return nil, err
	}

	return lastHash, nil

}
