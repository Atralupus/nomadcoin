package main

import (
	"github.com/Atralupus/nomadcoin/blockchain"
	"github.com/Atralupus/nomadcoin/cli"
	"github.com/Atralupus/nomadcoin/db"
)

func main() {
	defer db.Close()

	blockchain.Blockchain()
	cli.Start()
}
