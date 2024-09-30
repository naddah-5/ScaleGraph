package main

import (
	"fmt"
	"os"
	scaleGraph "scalegraph/src"
	"sync"
	"time"
)

func main() {
	alphaScript()
}

// There are three important parts to the loop:
// delay, done, and prt.
// delay is a block channel, when it is closed all spawned nodes will start their script.
// If startDelay != delay, the nodes will start as soon as they spawn.
// done is a waitgroup channel, when a node completes its script it notifies the main function over this  channel.
// prt is a blocking channel, when it is closed all nodes will log their termination.
func alphaScript() {
	fmt.Println("hello world")
	fmt.Printf("%+v\n", time.Now())

	var delay chan struct{}

	delay = make(chan struct{})
	done := make(chan struct{}, 100)
	prt := make(chan struct{})
	var wg sync.WaitGroup

	s := scaleGraph.NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	// max 10 000 nodes joining for now with logging
	// max 100 000 nodes joining for now with NO logging
	// limited number of goroutines?
	for i := 0; i < 5; i++ {
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
	time.Sleep(1*time.Second)
	close(prt)
	fmt.Println("closing")
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", time.Now())
	os.Exit(0)

}

func betaScript() {
	scaleGraph.TestFindNode()
	time.Sleep(5 * time.Second)
}
