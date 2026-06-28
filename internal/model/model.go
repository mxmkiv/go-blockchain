package model

import (
	"time"
)

// network has a virtual coin called Coin and a subcoin called Cent
// 1 Coin = 100 Cent
const CoinInCent = uint64(100)

type Cent uint64

// first 20 bytes of the public key hash (eth)
type Address [20]byte

// sha256 hash size
type HashType [32]byte

type HeaderBlock struct {
	PrevHash  HashType
	Nonce     uint64
	TimeStamp time.Time
	TxRoot    HashType
	StateRoot HashType
}

type Block struct {
	HeaderBlock
	BlockHash HashType
	Txs       []Transaction
}

type Transaction struct {
	TxNonce   uint64
	From      Address
	To        Address
	Amount    Cent
	Signature []byte // ed25519 signature
}
