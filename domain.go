package main

import "math/big"

// small example domain model
type BlockNumber *big.Int
type BlockHash string
type GasLimit *big.Int
type GasValue *big.Int
type Nonce uint64

type Block struct {
	Number   BlockNumber `json:"number"`
	Hash     BlockHash   `json:"hash"`
	GasLimit GasValue    `json:"gasLimit"`
	GasUsed  GasValue    `json:"gasUsed"`
	Nonce    Nonce       `json:"nonce"`
}

func NewBlockNumber(n int64) BlockNumber {
	return big.NewInt(n)
}
