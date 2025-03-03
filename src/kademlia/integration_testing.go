package kademlia

import (
	"fmt"
	"log"
	"slices"
	"time"
)

type result struct {
	nodeIndex    int
	missingTotal int
	missingIndex []int
}


func IntegrationTestFindNodeAny() bool {
	done := make(chan struct{}, 64)
	testName := "IntegrationTestFindNodeAny"
	verbose := false
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(50, done)
	<-done
	time.Sleep(time.Millisecond * 50)
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
		return false
	}
	return true
}

func IntegrationTestFindNodeSpecific() bool {
	done := make(chan struct{}, 64)
	testName := "IntegrationTestFindNodeSpecific"
	verbose := false
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(50, done)
	<-done
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

		return false
	}
	return true
}

func IntegrationTestMassiveFindNodeSpecific() bool {
	testName := "IntegrationTestMassiveFindNodeSpecific"
	verbose := false
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
		}
	}
	log.Printf("\nfailed %d of %d times", failCounter, mass)
	if failCounter > 0 {
		return false
	} else {
		return true
	}
}

func findNodeSpecific(verbose bool, testName string) bool {
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(500, done)
	<-done

	firstNode := nodes[0]
	lastNode := nodes[len(nodes)-1]
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

func IntegrationTestFindNodeVisibleNodes() bool {
	verbose := true
	testName := "IntegrationTestFindVisibleNodes"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	testSize := 100
	s := NewServer(false, 0.0)
	go s.StartServer()
	s.SpawnCluster(testSize, done)
	nodes := s.AllNodePointers()
	<-done

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
				lostNodes++
			}
		}
	}
	verPrint += fmt.Sprintf("failed to locate %d nodes", lostNodes)
	if verbose {
		fmt.Printf("\r\n")
		log.Printf(verPrint)
	}
	if lostNodes > 0 {
		return false
	} else {
		return true
	}
}

func IntegrationTestStoreAndFindAccount() bool {
	verbose := true
	testName := "IntegrationTestStoreAndFindAccount"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	testSize := 50
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := make([]*Node, 0, testSize)
	nodes = s.SpawnCluster(testSize, done)
	<-done

	time.Sleep(time.Millisecond * 100)

	if verbose {
		log.Println("Stimulating network...")
	}
	s.Stimulate()

	accID := RandomID()
	nodes[0].StoreAccount(accID)
	res, err := nodes[len(nodes)-1].FindAccount(accID)
	if verbose {
		verPrint += fmt.Sprintf("found account %v in nodes:\n", accID)
		for _, n := range res {
			verPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", n.ID(), RelativeDistance(n.ID(), accID))
		}
		log.Println(verPrint)
	}
	if err != nil {
		log.Println(err.Error())
		return false
	}

	nodeCon := make([]Contact, 0, len(nodes))
	for _, n := range nodes {
		nodeCon = append(nodeCon, n.Contact)
	}
	SortContactsByDistance(&nodeCon, accID)
	valPrint := fmt.Sprintf("the %d closest nodes to account %v in test:\n", REPLICATION, accID)
	for i := range REPLICATION {
		valPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", nodeCon[i].ID(), RelativeDistance(nodeCon[i].ID(), accID))
	}
	log.Println(valPrint)

	matches := "matching nodes\n"
	missMatch := 0
	for i, n := range res {
		if n.ID() != nodeCon[i].ID() {
			missMatch++
			matches += fmt.Sprintf("nodes at index %d do not match\n", i)
		} else {
			matches += fmt.Sprintf("nodes at index %d match\n", i)
		}
	}
	log.Println(matches)
	if missMatch != 0 {
		for i, n := range nodeCon {
			log.Printf("node %3d: %10v\n", i, n.ID())
		}
		return false
	}

	return true
}

func IntegrationTestStoreAndFindAccountFromSharingID() bool {
	verbose := true
	testName := "IntegrationTestStoreAndFindAccountFromSharingID"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	testSize := 100
	s := NewServer(false, 0.0)
	go s.StartServer()
	nodes := s.SpawnCluster(testSize, done)
	<-done

	if verbose {
		log.Printf("Stimulating network...\n")
	}
	s.Stimulate()
	time.Sleep(time.Millisecond * 100)

	accID := nodes[0].ID()
	nodes[0].StoreAccount(accID)
	res, err := nodes[len(nodes)-1].FindAccount(accID)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	nodeCon := make([]Contact, len(nodes))
	for _, n := range nodes {
		nodeCon = append(nodeCon, n.Contact)
	}
	SortContactsByDistance(&nodeCon, accID)
	valPrint := fmt.Sprintf("the %d closest nodes to account %v in test:\n", REPLICATION, accID)
	for i := range min(REPLICATION, len(nodeCon)) {
		valPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", nodeCon[i].ID(), RelativeDistance(nodeCon[i].ID(), accID))
	}

	matches := "matching nodes\n"
	missMatch := false
	for i, n := range res {
		if n.ID() != nodeCon[i].ID() {
			matches += fmt.Sprintf("nodes at index %d do not match\n", i)
			missMatch = true
		} else {
			matches += fmt.Sprintf("nodes at index %d match\n", i)
		}
	}
	if missMatch {
		if verbose {
			verPrint += fmt.Sprintf("found account %v in nodes:\n", accID)
			for _, n := range res {
				verPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", n.ID(), RelativeDistance(n.ID(), accID))
			}
			log.Println(verPrint)
		}
		log.Println(valPrint)
		log.Println(matches)
	}

	return true
}

func IntegrationTestStoreAndFindAccountFromAllNodes() bool {
	verbose := true
	testName := "IntegrationTestStoreAndFindAccountFromSharingID"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	inspectRoutingTables := false
	testSize := 200
	stimulate := 1
	s := NewServer(false, 0.0)
	go s.StartServer()
	s.SpawnCluster(testSize, done)
	<-done

	if verbose {
		log.Printf("Stimulating network...\n")
	}

	for range stimulate {
		err := s.Stimulate()
		if err != nil {
			log.Println(err.Error())
		}
	}

	nodes := s.AllNodePointers()
	accID := RandomID()
	tmp := make([]Contact, 0, len(nodes))
	for _, n := range nodes {
		tmp = append(tmp, n.Contact)
	}
	SortContactsByDistance(&tmp, accID)
	nodeCon := make([]Contact, 0, REPLICATION)
	for i := range REPLICATION {
		nodeCon = append(nodeCon, tmp[i])
	}
	nodes[0].StoreAccount(accID)
	failFinds := make([]int, len(nodes))
	respChan := make(chan result, 128)
	time.Sleep(time.Second)
	for i, origin := range nodes {
		time.Sleep(time.Millisecond * 10)
		go func(respChan chan result, i int, origin *Node) {
			//fmt.Printf("\rsearching from node %3d, %10v", i, origin.ID())
			res, _ := origin.FindAccount(accID)
			missing := 0
			if len(res) > len(nodeCon) {
				fmt.Printf("wtf\n")
			}
			for _, con := range res {
				if !slices.Contains(nodeCon, con) {
					missing++
				}
			}
			respChan <- result{i, missing, make([]int, 0)}
		}(respChan, i, origin)

	}
	for range nodes {
		res := <-respChan
		failFinds[res.nodeIndex] = res.missingTotal
	}
	verPrint += fmt.Sprintf("incorrect validators for the following nodes:\n")
	for i, missing := range failFinds {
		if missing != 0 {
			verPrint += fmt.Sprintf("node %3d - %10v - missing %3d validators, distance from acc - %10v\n", i, nodes[i].ID(), missing, RelativeDistance(nodes[i].ID(), accID))
			if inspectRoutingTables {
				for _, n := range nodes {
					if n.ID() == nodes[i].ID() {
						verPrint += fmt.Sprintf("Routing table for node %10d\n", n.ID())
						verPrint += fmt.Sprintf("%s", n.Display())
						break
					}
				}
			}
		}
	}

	valPrint := fmt.Sprintf("the %d closest nodes to account %v in test:\n", REPLICATION, accID)
	for _, con := range nodeCon {
		valPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", con.ID(), RelativeDistance(con.ID(), accID))
	}

	log.Println(valPrint)
	log.Println(verPrint)

	return true
}

func IntegrationTestStoreAndDisplayAccount() bool {
	testName := "IntegrationTestStoreAndDisplayAccount"
	done := make(chan struct{}, 64)
	testSize := 100
	stimulation := 1
	s := NewServer(false, 0.0)
	go s.StartServer()
	s.SpawnCluster(testSize, done)
	<-done
	for range stimulation {
		log.Println("stimualting network...")
		s.Stimulate()
		time.Sleep(TIMEOUT)
	}

	nodes := s.AllNodePointers()
	accID := RandomID()
	nodes[0].InsertAccount(accID)
	time.Sleep(time.Second)
	log.Printf("storing account %10v", accID)

	displayString, err := nodes[1].DisplayAccount(accID)
	if err != nil {
		log.Printf("%s [ERROR] - %s\n", testName, err.Error())
	} else {
		log.Printf("%s\n%s\n", testName, displayString)
	}

	for _, n := range nodes {
		if n.scalegraph.StoredAccountCount() != 0 {
			for _, stored := range n.scalegraph.StoredAccounts() {
				log.Printf("node %10v storing account %10v", n.ID(), stored)
			}
		}
	}

	return true
}

func IntegrationTestStoreAndLockAccount() bool {
	testName := "IntegrationTestStoreAndLockAccount"
	log.Println(testName)
	done := make(chan struct{}, 64)
	testSize := 50
	stimulation := 1
	s := NewServer(false, 0.0)
	go s.StartServer()
	s.SpawnCluster(testSize, done)
	<-done
	for range stimulation {
		log.Println("stimualting network...")
		s.Stimulate()
		time.Sleep(TIMEOUT)
	}

	nodes := s.AllNodePointers()
	accID := RandomID()
	nodes[0].InsertAccount(accID)
	log.Printf("storing account %10v", accID)

	return true
}

func IntegrationNodeLookupDataGathering(testSize int) (int, []int) {
	verbose := true
	testName := "IntegrationTestStoreAndFindAccountFromSharingID"
	done := make(chan struct{}, 64)
	verPrint := fmt.Sprintf("[%s]\n", testName)
	stimulate := 1
	s := NewServer(false, 0.0)
	go s.StartServer()
	s.SpawnCluster(testSize, done)
	<-done

	if verbose && stimulate > 0 {
		log.Printf("Stimulating network...\n")
	}

	for range stimulate {
		s.Stimulate()
	}

	nodes := s.AllNodePointers()
	missingValidators := make([][]int, len(nodes))
	accID := RandomID()
	tmp := make([]Contact, 0, len(nodes))
	for _, n := range nodes {
		tmp = append(tmp, n.Contact)
	}
	SortContactsByDistance(&tmp, accID)
	nodeCon := make([]Contact, 0, REPLICATION)
	for i := range REPLICATION {
		nodeCon = append(nodeCon, tmp[i])
	}
	nodes[0].StoreAccount(accID)
	respChan := make(chan result, 128)
	time.Sleep(time.Second)
	for i, origin := range nodes {
		time.Sleep(time.Millisecond * 10)
		go func(respChan chan result, i int, origin *Node) {
			res, _ := origin.FindAccount(accID)
			missingIndecies := make([]int, 0)
			for i, con := range res {
				if !slices.Contains(nodeCon, con) {
					missingIndecies = append(missingIndecies, i)
				}
			}
			respChan <- result{i, 1, missingIndecies}
		}(respChan, i, origin)

	}
	for range nodes {
		res := <-respChan
		missingValidators[res.nodeIndex] = res.missingIndex
	}
	verPrint += fmt.Sprintf("incorrect validators for the following nodes:\n")
	missingAVal := 0
	for i, missing := range missingValidators {
		if len(missing) != 0 {
			missingAVal++
			verPrint += fmt.Sprintf("node %3d - %10v - missing %3d validators, indecies: %v\n", i, nodes[i].ID(), len(missing), missing)
		}
	}

	//valPrint := fmt.Sprintf("the %d closest nodes to account %v in test:\n", REPLICATION, accID)
	// for _, con := range nodeCon {
	// 	valPrint += fmt.Sprintf("node: %10v, distance from account: %10v\n", con.ID(), RelativeDistance(con.ID(), accID))
	// }

	//log.Println(valPrint)
	//log.Println(verPrint)

	testRes := make([]int, REPLICATION)

	for _, datum := range missingValidators {
		for _, missing := range datum {
			testRes[missing]++
		}
	}


	return missingAVal, testRes
}
