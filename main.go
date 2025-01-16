package main

import (
	"fmt"
	"log"
	"main/src/kademlia"
)

func main() {
	res := kademlia.IntegrationTestFindNodeVisibleNodes()
	if !res {
		log.Println("[TESTING FAILED]")
	}
}

func script() {
	done := make(chan [5]uint32, 64)
	s := kademlia.NewServer(true, 0.0)
	masterNode := s.MasterNode()
	go s.StartServer()
	var nodes []*kademlia.Node
	for i := 0; i < 5; i++ {
		node := s.SpawnNode(done)
		nodes = append(nodes, node)
	}
	nodes = append(nodes, s.SpawnNode(done))
	debugNode := nodes[len(nodes)-1]
	debugNode.Debug(true)
	// debugNode := masterNode
	// time.Sleep(time.Millisecond * 15000)
	for i := range nodes {
		cID := <-done
		log.Printf("%v finished, %d channels done", cID, i+1)
	}
	nodeState := ""
	for _, val := range nodes {
		if val.ID() == debugNode.ID() {
			continue
		}
		nodeState += fmt.Sprintf("\n%s\n", val.Display())
	}
	nodeState += fmt.Sprintf("\nDebug Node:\n%s\n", debugNode.Display())
	nodeState += fmt.Sprintf("\nMaster Node:\n%s", masterNode.Display())
	log.Println(nodeState)
	return
}
