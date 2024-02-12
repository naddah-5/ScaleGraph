package scalegraph

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	addr  net.UDPAddr
	conn  net.Conn
}

func NewServer(addr [4]byte) *Server {
	log.Println("starting a new server")
	var ip net.IP
	for i := 0; i < 4; i++ {
		ip = append(ip, addr[i])
	}
	fmtAddr := fmt.Sprintf("127.0.0.1:%d", PORT)
	udpAddr, err := net.ResolveUDPAddr("udp", fmtAddr)
	if err != nil {
		log.Println(err)
	}

	return &Server{
		addr:  *udpAddr,
	}
}

func (s *Server) read() {
	buf := make([]byte, 2048)

	for {
		n, err := s.conn.Read(buf)
		if err != nil {
			log.Println("read error: ", err)
		} else {
			msg := buf[:n]
			log.Printf(string(msg))
			// make a go call to the handler
		}
	}
}
