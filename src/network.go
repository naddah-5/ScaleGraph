package scalegraph


import (
	"log"
	"net"
)

func TestVisibility() {
	log.Println("hello world")
}

type Server struct {
	addr  net.IP
	ln    net.Listener
	close chan struct{}
}

func NewServer(addr [4]byte) *Server {
	log.Println("starting a new server")
	var ip net.IP
	for i := 0; i < 4; i++ {
		ip = append(ip, addr[i])
	}

	return &Server{
		addr:  ip,
		close: make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr.String())
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.accept()

	<-s.close

	return nil
}

func (s *Server) accept() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			continue
		}

		go s.read(conn)
		log.Printf("accepted connection from: %s", conn.RemoteAddr().String())
	}
}

func (s *Server) read(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("read error: ", err)
		} else {
			msg := buf[:n]
			log.Printf(string(msg))
		}
	}
}
