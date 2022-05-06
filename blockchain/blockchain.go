package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	data     string
	hash     string
	prevHash string
}

func (b *Block) calcBlockHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	b.hash = hexHash
}

type blockChain struct {
	blocks []*Block
}

var b *blockChain
var once sync.Once

func GetBlockChain() *blockChain {
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
	return b.blocks[b.Size()-1].hash
}

func (b blockChain) Size() int {
	return len(b.blocks)
}

func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b blockChain) PrintBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s, Hash: %s, PrevHash: %s\n", block.data, block.hash, block.prevHash)
	}
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", b.getLastBlockHash()}
	newBlock.calcBlockHash()

	return &newBlock
}
