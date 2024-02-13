package main

import (
	"log"
	scaleGraph "scalegraph/src"
	"sync"
	"time"
)



func main() {
	var wg sync.WaitGroup
	log.Println("hello world")
	s := scaleGraph.NewServer([4]byte{127, 0, 0, 1})
	if s == nil {
		log.Fatal("server could not be created")
	}
	wg.Add(1)
	go s.Start()
	time.Sleep(2 * time.Second)
	s.Close()
	wg.Done()
	wg.Wait()
}
