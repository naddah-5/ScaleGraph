package main

import (
	"fmt"
	"log"
	"os"
	scaleGraph "scalegraph/src"
	"sync"
	"time"
)

// There are three important parts to the loop:
// delay, done, and prt.
// delay is a block channel, when it is closed all spawned nodes will start their script.
// If startDelay != delay, the nodes will start as soon as they spawn.
// done is a waitgroup channel, when a node completes its script it notifies the main function over this  channel.
// prt is a blocking channel, when it is closed all nodes will log their termination.
func main() {
	fmt.Println("hello world")
	fmt.Printf("%+v\n", time.Now())

	delay := make(chan struct{})
	done := make(chan struct{}, 100)
	prt := make(chan struct{})
	var wg sync.WaitGroup

	s := scaleGraph.NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	for i := 1; i < 10; i++ {
		var startDelay chan struct{} = nil
		if startDelay == nil {
			time.Sleep(100 * time.Millisecond)
		}
		wg.Add(1)
		s.SpawnNode(startDelay, done, prt)
	}
	close(delay)

	// handle the done channel
	go func(done chan struct{}, wg *sync.WaitGroup) {
		for {
			<-done
			log.Println("received done")
			wg.Done()
		}
	}(done, &wg)

	wg.Wait()
	close(prt)
	time.Sleep(15 * time.Second)
	fmt.Printf("%+v\n", time.Now())
	os.Exit(0)
}
