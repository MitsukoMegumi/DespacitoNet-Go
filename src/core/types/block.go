package types

import (
	"github.com/mitsukomegumi/Despacito-Go/src/common"
)

// Block - blockchain block
type Block struct {
	Reward       int    `json:"reward"`
	MinerAddress string `json:"minerwallet"`

	Version int `json:"version"`

	DespacitoSrc *[]byte `json:"despacito"`

	Transactions *[]Transaction `json:"transactions"`

	Hash string `json:"hash"`
}

// NewBlock - creates instance of block struct
func NewBlock(reward int, miner string, despacito *[]byte, version int, transactions *[]Transaction) (*Block, error) {
	blck := Block{Reward: reward, MinerAddress: miner, Version: version, DespacitoSrc: despacito, Transactions: transactions, Hash: ""}
	hash, err := common.SHA256(blck)

	if err != nil {
		return nil, err
	}

	blck.Hash = hash

	return &blck, nil
}