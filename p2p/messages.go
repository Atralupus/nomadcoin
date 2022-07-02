package p2p

import (
	"encoding/json"

	"github.com/Atralupus/nomadcoin/blockchain"
	"github.com/Atralupus/nomadcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksReponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func (m *Message) addPayload(payload interface{}) {
	b, err := json.Marshal(payload)
	utils.HandleErr(err)
	m.Payload = b
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}
	m.addPayload(payload)
	mJson, err := json.Marshal(m)
	utils.HandleErr(err)
	return mJson
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}
