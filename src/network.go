package scalegraph

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	addr  net.UDPAddr
	conn  net.UDPConn
	close chan struct{}
}

func NewServer(addr [4]byte) *Server {
	log.Println("creating a new server")
	var ip net.IP
	for i := 0; i < 4; i++ {
		ip = append(ip, addr[i])
	}
	fmtAddr := fmt.Sprintf("127.0.0.1:%d", PORT)
	udpAddr, err := net.ResolveUDPAddr("udp", fmtAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		addr: *udpAddr,
		conn: *conn,
		close: make(chan struct{}),
	}
}

func (s *Server) Start() {
	defer s.conn.Close()
	go s.listen()
	log.Println("starting a server")
	<-s.close
	return
}

func (s *Server) Close() {
	log.Println("closing server")
	close(s.close)
	log.Println("returning")
	return
}

func (s *Server) listen() {
	for {
		buf := make([]byte, 2048)
		n, _, err := s.conn.ReadFrom(buf)
		if err != nil {
			log.Println("read error: ", err)
		} else {
			msg := buf[:n]
			log.Printf(string(msg))
			go Handler(msg)
		}
	}

}
