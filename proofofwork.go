package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// targetBits define complexity for mining
const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, uint(256-targetBits))

	return &ProofOfWork{b, t}
}

func (pow *ProofOfWork) prepare(nonce int) []byte {

	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			int64ToHex(pow.block.Timestamp),
			int64ToHex(int64(targetBits)),
			int64ToHex(int64(nonce)),
		},
		[]byte{},
	)

}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepare(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1

}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte

	// Init nonce, in cryptography, a nonce is an arbitrary number that can only be used once. like a counter
	// TODO: Read more https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonce := 0

	fmt.Printf("Mining the block with data \"%s\"\n", pow.block.Data)
	// TODO: Search if that there is another algorithm for this.
	for nonce < math.MaxInt64 {
		data := pow.prepare(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Utils
// IntToHex converts an int64 to a byte array
func int64ToHex(n int64) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, n)
	if err != nil {
		log.Panic(err)
	}

	return b.Bytes()
}
