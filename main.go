package main

import (
	"log"
	"main/src/kademlia"
	"time"
)


func main() {
	script()
}

func script() {
	s := kademlia.NewServer()
	go s.StartServer()
	var nodes []*kademlia.Node
	for i := 0; i < 2; i++ {
		node := s.SpawnNode()
		nodes = append(nodes, node)
	}
	time.Sleep(time.Millisecond * 5000)
	for _, val := range nodes {
		log.Print(val.Display())
	}
	return
}
