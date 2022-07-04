package p2p

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Atralupus/nomadcoin/blockchain"
	"github.com/Atralupus/nomadcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksReponse
	MessageNewBlockNotify
	MessageNewTxNotify
	MessageNewPeerNotify
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJson(payload),
	}

	return utils.ToJson(m)
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksReponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func notifyNewBlock(b *blockchain.Block, p *peer) {
	m := makeMessage(MessageNewBlockNotify, p)
	p.inbox <- m
}

func notifyNewTx(tx *blockchain.Tx, p *peer) {
	m := makeMessage(MessageNewTxNotify, p)
	p.inbox <- m
}

func notifyNewPeer(addtress string, p *peer) {
	m := makeMessage(MessageNewPeerNotify, p)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	fmt.Printf("Peer: %s, Sent a meesage with kind of %d", p.key, m.Kind)
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		json.Unmarshal(m.Payload, &payload)

		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)

		if payload.Height >= b.Height {
			requestAllBlocks(p)
		} else {
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		sendAllBlocks(p)
	case MessageAllBlocksReponse:
		var payload []*blockchain.Block
		json.Unmarshal(m.Payload, &payload)
		blockchain.Blockchain().Replace(payload)
	case MessageNewBlockNotify:
		var payload *blockchain.Block
		json.Unmarshal(m.Payload, &payload)
		blockchain.Blockchain().AddPeerBlock(payload)
	case MessageNewTxNotify:
		var payload *blockchain.Tx
		json.Unmarshal(m.Payload, &payload)
		blockchain.Mempool().AddPeerTx(payload)
	case MessageNewPeerNotify:
		var payload string
		json.Unmarshal(m.Payload, &payload)
		parts := strings.Split(payload, ":")
		AddPeer(parts[0], parts[1], parts[2], false)
	}
}
