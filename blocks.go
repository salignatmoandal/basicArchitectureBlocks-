package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block headers
type Block struct {
	Timestamps    int64  // is the current timestamp (when the block is created)
	Data          []byte // is the actual valuable information containing in the block
	PrevBlockHash []byte //store the hash of the previous block, and hash is the hash of the block.
	Hash          []byte
}

// Create a function who allow to Set Hash calculate and sets bock hash

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamps, 10))                      // Convertir le timestamps en un tableau de bytes
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // Combiner les données nécessaire pour calculer le hash
	hash := sha256.Sum256(headers)                                                // Calculer le hachage de l'en-tête

	b.Hash = hash[:] // Assigner le hash calculé à la proprieté hash du bloc
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
