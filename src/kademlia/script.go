package kademlia

import (
	"fmt"
	"sync"
	"time"
)

func AlphaScript() {
	fmt.Println("Starting Alpha Script")

	delay := make(chan struct{})
	done := make(chan struct{})
	prt := make(chan struct{})

	var wg sync.WaitGroup

	s := NewServer()
	go s.StartServer()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		node := s.SpawnNode()
		go node.Start()
	}
}
