package main

type Block struct {
	Timestamps    int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
