package main

import "github.com/Atralupus/nomadcoin/blockchain"

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	chain.PrintBlocks()
}
