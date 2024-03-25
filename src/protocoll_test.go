package scalegraph

import (
	"log"
	"testing"
)

func generateTestNodes(x int, serverIP [4]byte, nodeListener chan RPC, nodeSender chan RPC, masterIP [4]byte) []*Node {
	nodes := make([]*Node, x)
	for i := range x {
		id := [5]uint32{0, 0, 0, 0, 0}
		id[4] += uint32(i)
		ip := [4]byte{0, 0, 0, 1}
		ip[3] += byte(i)
		node := NewNode(id, ip, nodeListener, nodeSender, serverIP, masterIP)
		nodes[i] = &node
	}
	
	return nodes
}

func testingServer(nodes int) *Simnet {
	s := NewServer()
	go s.StartServer()

	spawn := generateTestNodes(nodes, s.serverIP, s.listener, make(chan RPC), s.masterNode)
	
	for _, node := range spawn {
		s.AttachThisNode(node)
	}
	return s
}

func TestFindNode(t *testing.T) {
	s := testingServer(10)
	state := s.AllNodes()
	for _, nodemap := range state {
		log.Printf("%+v", nodemap)
	}


}
