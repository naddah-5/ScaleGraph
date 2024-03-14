package main

import (
	"fmt"
	"os"
	scaleGraph "scalegraph/src"
	"sync"
	"time"
)

// There are three important parts to the loop:
// delay, done, and prt.
// delay is a block channel, when it is closed all spawned nodes will start their script.
// done is a waitgroup channel, when a node completes its script it notifies the main function over this  channel.
// prt is a blocking channel, when it is closed all nodes will log their termination.
func main() {
	fmt.Println("hello world")
	fmt.Printf("%+v\n", time.Now())

	delay := make(chan struct{}, 100)
	done := make(chan struct{})
	prt := make(chan struct{})
	var wg sync.WaitGroup

	s := scaleGraph.NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	for i := 1; i < 100; i++ {
		wg.Add(1)
		s.SpawnNode(delay, done, prt)
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
	time.Sleep(15 * time.Second)
	os.Exit(0)
}
