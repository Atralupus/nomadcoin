package p2p

import (
	"fmt"
	"net/http"

	"github.com/Atralupus/nomadcoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	openPort := r.URL.Query().Get("openPort")
	// ip := utils.Splitter(r.RemoteAddr, ":", 0)
	// upgrader.CheckOrigin = func(r *http.Request) bool {
	// 	return openPort != "" && ip != ""
	// }
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	peer := initPeer(conn, "localhost", openPort)
	peer.inbox <- []byte("Hello.")
}

func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
