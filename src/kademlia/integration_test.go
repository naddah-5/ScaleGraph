package kademlia

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestFindNodeAny(t *testing.T) {
	done := make(chan [5]uint32, 64)
	testName := "TestFindNodeAny"
	verbose := false
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	var nodes []*Node
	for i := 0; i < 5; i++ {
		node := s.SpawnNode(done)
		nodes = append(nodes, node)
	}
	for range nodes {
		<-done
	}
	time.Sleep(time.Millisecond * 1000)
	firstNode := nodes[0]
	lastNode := nodes[len(nodes)-1]
	if verbose {
		verPrint += fmt.Sprintf("Looking for %v in simulation state:\n", firstNode.ID())
		verPrint += "==============================================================="
		for _, node := range nodes {
			verPrint += fmt.Sprintf("\n%s\n", node.Display())
		}
		verPrint += "===============================================================\n"
	}
	res := firstNode.FindNode(lastNode.ID())
	if verbose {
		verPrint += fmt.Sprintf("Looking for %v\n", lastNode.ID())
		verPrint += fmt.Sprintf("Found nodes\n")
		for _, rNode := range res {
			verPrint += fmt.Sprintf("%v\n", rNode.Display())
		}
	}

	if verbose {
		log.Printf(verPrint)
	}
	if len(res) == 0 {
		log.Printf("[%s] - failed to find any nodes at all", testName)
		log.Printf("res length: %d\nres contains:\n", len(res))
		for _, n := range res {
			log.Printf("%s\n", n.Display())
		}
		t.Fail()
	}
}

func TestFindNodeSpecific(t *testing.T) {
	done := make(chan [5]uint32, 64)
	testName := "TestFindNodeSpecific"
	verbose := false
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	var nodes []*Node
	for range 50 {
		node := s.SpawnNode(done)
		nodes = append(nodes, node)
	}
	for range nodes {
		<-done
	}
	firstNode := nodes[0]
	lastNode := nodes[len(nodes)-1]
	time.Sleep(time.Second * 1)
	res := firstNode.FindNode(lastNode.ID())

	if verbose {
		log.Printf(verPrint)
	}
	if res[0].ID() != lastNode.ID() {
		log.Printf("[%s] - head result does not match\nlooking for: %v", testName, lastNode.ID())
		log.Printf("res length: %d\nres contains:\n", len(res))
		for _, n := range res {
			log.Printf("%s\n", n.Display())
		}

		t.Fail()
	}
}

func TestMassiveFindNodeSpecific(t *testing.T) {
	passing := true
	if passing {
		return
	}
	testName := "TestMassiveFindNodeSpecific"
	verbose := true
	if verbose {
		log.Printf("[%s] - starting test\n", testName)
	}

	mass := 10
	failCounter := 0
	for i := range mass {
		itterRes := findNodeSpecific(verbose, testName)
		if !itterRes {
			failCounter++
			log.Printf("failed %d of %d - %d times", failCounter, i+1, mass)
			t.Fail()
		}
	}
	log.Printf("\nfailed %d of %d times", failCounter, mass)
}

func findNodeSpecific(verbose bool, testName string) bool {
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(5000, done)
	<-done

	firstNode := nodes[0]
	lastNode := nodes[len(nodes)-1]
	time.Sleep(time.Millisecond * 500)
	res := firstNode.FindNode(lastNode.ID())
	if verbose {
		verPrint += fmt.Sprintf("Looking for %v\n", lastNode.ID())
		verPrint += fmt.Sprintf("Found nodes\n")
		for _, rNode := range res {
			verPrint += fmt.Sprintf("%v\n", rNode.Display())
		}
		log.Println(verPrint)
	}

	if !SliceContains(lastNode.ID(), &res) {
		if verbose {
			verPrint += fmt.Sprintf("[%s] - head result does not match", testName)
			verPrint += fmt.Sprintf("looking for: %v\n", lastNode.ID())
			verPrint += fmt.Sprintf("res length: %d\nres contains:\n", len(res))
			for _, n := range res {
				verPrint += fmt.Sprintf("%s\n", n.Display())
			}

			log.Println(verPrint)
		}
		return false
	}
	return true
}

func TestFindNodeVisibleNodes(t *testing.T) {
	passing := false
	if passing {
		return
	}
	verbose := true
	testName := "TestFindVisibleNodes"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	testSize := 1000
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(testSize, done)
	<-done

	time.Sleep(time.Millisecond * 100)
	lostNodes := 0
	for i, origin := range nodes {
		if verbose {
			fmt.Printf("\rSearching from node %d", i)
		}
		for _, node := range nodes {
			res := origin.FindNode(node.ID())
			if res[0].ID() != node.ID() {
				lostNodes++
				verPrint += fmt.Sprintf("Failed to locate node: %v\n", node.ID())
				t.Fail()
			}
		}
	}
	verPrint += fmt.Sprintf("failed to locate %d nodes", lostNodes)
	if verbose {
		log.Printf(verPrint)
	}
}
