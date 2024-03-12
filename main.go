package main

import (
	"fmt"
	"log"
	"os"
	scaleGraph "scalegraph/src"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("%+v\n", time.Now())

	s := scaleGraph.NewServer()
	go s.StartServer()
	time.Sleep(1 * time.Second)
	for i := 1; i < 500; i++ {
		s.SpawnNode()
	}

	time.Sleep(200 * time.Second)
	nodes := s.AllNodes()
	currentNodes := ""
	currentNodes += fmt.Sprintf("all current nodes:\n")
	for _, n := range nodes {
		currentNodes += fmt.Sprintf("%+v\n", n)
	}
	log.Println(currentNodes)
	os.Exit(0)
}
