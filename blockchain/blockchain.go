package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
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

func createBlock(data string) *Block {
	newBlock := Block{data, "", b.getLastBlockHash()}
	newBlock.calcBlockHash()

	return &newBlock
}
