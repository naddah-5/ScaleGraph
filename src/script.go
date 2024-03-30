package scalegraph

import (
	"fmt"
	"log"
	"time"
)

// Generates x number of test nodes with incrementing id's and ip's.
func generateTestNodes(x byte, serverIP [4]byte, nodeListener chan RPC, nodeSender chan RPC, masterIP [4]byte) []*Node {
	nodes := make([]*Node, x)
	for i := range x {
		id := [5]uint32{0, 0, 0, 0, 0}
		id[4] += uint32(i)
		ip := [4]byte{0, 0, 0, 0}
		ip[3] += i
		node := NewNode(id, ip, nodeListener, nodeSender, serverIP, masterIP)
		nodes[i] = &node
	}

	return nodes
}

// Creates and returns a simulation server for testing purposes.
func testingServer(nodes byte) *Simnet {
	s := NewServer()
	go s.StartServer()

	spawn := generateTestNodes(nodes, s.serverIP, s.listener, make(chan RPC), s.masterNode)

	for _, node := range spawn {
		s.AttachThisNode(node)
	}
	return s
}

func TestFindNode() {
	s := testingServer(10)
	state := s.AllNodes()
	for _, nodemap := range state {
		log.Printf("%+v", nodemap)
	}
	ping := GenerateRPC(SEND, state[0], state[1].IP())
	s.listener <- ping
	s.listener <- ping
	fmt.Println("done")
}

func (node *Node) NodeAlphaScript(delay chan struct{}, done chan struct{}, prt chan struct{}) {
	if delay != nil {
		<-delay
	}

	node.Ping(node.serverIP)

	time.Sleep(1 * time.Second)
	go node.FindNode(node.ID())
	time.Sleep(10 * time.Second)

	comp := make(chan struct{})
	go func(node *Node, comp chan struct{}) {
		node.FindNode(node.ID())
		close(comp)
	}(node, comp)
	<-comp

	if done != nil {
		done <- struct{}{}
	}

	if prt != nil {
		_, allDone := <-prt
		if !allDone {
			if DEBUG {
				dumpTable := ""
				dumpTable += fmt.Sprintf("[node] - %+v - current routing table:\n", node.ID())
				for i := range node.router {
					node.router[i].lock.RLock()
					for c := node.router[i].content.Front(); c != nil; c = c.Next() {
						dumpTable += fmt.Sprintf("\tcontact: %+v\n", c.Value.(contact).id)
					}
					node.router[i].lock.RUnlock()
				}
				log.Println(dumpTable)
			}
		}
	}

}
