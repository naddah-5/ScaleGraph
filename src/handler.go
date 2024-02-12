package scalegraph

import "net"

type call struct {
	conn     net.Conn
	resChan  chan string
	CMD
}

func Handler(conn net.Conn) {
	
}
