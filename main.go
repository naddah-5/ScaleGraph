package main

import (
	"log"
	scaleGraph "scalegraph/src"
	"sync"
	"time"
)



func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Println("hello world")
	s := scaleGraph.NewServer()
	go s.StartServer()
	s.SpawnNode()
	nodes := s.AllNodes()
	time.Sleep(3 * time.Second)
	log.Println("all current nodes")
	for _, n := range(nodes) {
		log.Printf("%+v\n", n)
	}
	wg.Done()
}
