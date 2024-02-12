package scalegraph

import "net"

type call struct {
	conn    net.Conn
	resChan chan string
	RPC
}

func Handler(conn net.Conn) {

}
