package scalegraph

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
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

// Creates and returns a simulation server
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


// There are three important parts to the loop:
// delay, done, and prt.
// delay is a block channel, when it is closed all spawned nodes will start their script.
// If startDelay != delay, the nodes will start as soon as they spawn.
// done is a waitgroup channel, when a node completes its script it notifies the main function over this  channel.
// prt is a blocking channel, when it is closed all nodes will log their termination.
func AlphaScript() {
	fmt.Println("hello world")
	fmt.Printf("%+v\n", time.Now())

	var delay chan struct{}

	delay = make(chan struct{})
	done := make(chan struct{}, 100)
	prt := make(chan struct{})
	var wg sync.WaitGroup

	s := NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	// max 10 000 nodes joining for now with logging
	// max 100 000 nodes joining for now with NO logging
	// limited number of goroutines?
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		node := s.SpawnNode()
		go node.NodeAlphaScript(delay, done, prt)
	}
	close(delay)

	// handle the done channel
	go func(done chan struct{}, wg *sync.WaitGroup) {
		for {
			<-done
			wg.Done()
		}
	}(done, &wg)

	wg.Wait()
	close(prt)
	fmt.Println("closing")
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", time.Now())
	os.Exit(0)

}

func (node *Node) NodeAlphaScript(delay chan struct{}, done chan struct{}, prt chan struct{}) {
	if delay != nil {
		<-delay
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	go node.Ping(node.serverIP)

	time.Sleep(3 * time.Second)
	comp := make(chan struct{})
	go func(node *Node, comp chan struct{}) {
		node.FindNode(node.ID())

		time.Sleep(100 * time.Microsecond)

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
					for c := node.router[i].content.Front(); c != nil; c = c.Next() {
						dumpTable += fmt.Sprintf("\tcontact: %+v\n", c.Value.(contact).id)
					}
				}
				log.Println(dumpTable)
			}
		}
	}

}

// Script for passive node
func BetaScript(size int) {
	s := NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	nodes := make([]*Node, 0, size)

	for i := 0; i < size; i++ {
		nodes = append(nodes, s.SpawnNode())
	}

	time.Sleep(1 * time.Second)
	//----------
	// do stuff here
	wallet := NewWallet(GenerateID(), 0)

	entry := nodes[rand.Intn(len(nodes))]
	err := entry.StoreWallet(wallet)
	if err != nil {
		log.Println(err.Error())
	}

	res, err := entry.ShowWallet(wallet.walletID)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)

	//----------
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", time.Now())
	os.Exit(0)

}
