package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) calcBlockHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	hexHash := fmt.Sprintf("%x", hash)
	b.Hash = hexHash
}

type blockChain struct {
	blocks []*Block
}

var b *blockChain
var once sync.Once

func GetBlockchain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.AddBlock("Genesis Block")
		})
		return b
	}
	return b
}

func (b blockChain) getLastBlockHash() string {
	if b.Size() == 0 {
		return ""
	}
	return b.blocks[b.Size()-1].Hash
}

func (b blockChain) Size() int {
	return len(b.blocks)
}

func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockChain) AllBlocks() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("block not found")

func (b *blockChain) GetBlock(height int) (*Block, error) {
	if height > b.Size() {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", b.getLastBlockHash(), b.Size() + 1}
	newBlock.calcBlockHash()

	return &newBlock
}
