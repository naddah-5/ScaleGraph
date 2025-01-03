package kademlia

import (
	"log"
	"testing"
)

func TestRoutingTableFindXClosestOrder(t *testing.T) {
	testName := "TestRoutingTableFindXClosestOrder"
	verbose := false
	router := NewRoutingTable([5]uint32{0, 0, 0, 0, 0}, 160, 10)
	for i := 0; i < 10000; i++ {
		router.AddContact(NewRandomContact())
	}
	nodeA := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	nodeB := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 2})
	nodeC := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 3})
	nodeD := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 4})
	nodeE := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 5})
	nodeF := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 6})
	nodeG := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 7})
	nodeH := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 8})
	nodeI := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 9})
	nodeJ := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 10})

	router.AddContact(nodeA)
	router.AddContact(nodeB)
	router.AddContact(nodeC)
	router.AddContact(nodeD)
	router.AddContact(nodeE)
	router.AddContact(nodeF)
	router.AddContact(nodeG)
	router.AddContact(nodeH)
	router.AddContact(nodeI)
	router.AddContact(nodeJ)

	res, err := router.FindXClosest(10, [5]uint32{0, 0, 0, 0, 0})
	if err != nil {
		log.Printf("[%s] - failed... %s", testName, err.Error())
		t.Fail()
	}
	if verbose {
		log.Printf("[%s] - contacts in node", testName)
		for _, v := range res {
			log.Printf("contact: %10v", v)
		}
	}
	if res[0] != nodeA {
		t.Fail()
	}
	if res[1] != nodeB {
		t.Fail()
	}
	if res[2] != nodeC {
		t.Fail()
	}
	if res[3] != nodeD {
		t.Fail()
	}
	if res[4] != nodeE {
		t.Fail()
	}
	if res[5] != nodeF {
		t.Fail()
	}
	if res[6] != nodeG {
		t.Fail()
	}
	if res[7] != nodeH {
		t.Fail()
	}
	if res[8] != nodeI {
		t.Fail()
	}
	if res[9] != nodeJ {
		t.Fail()
	}

}

func TestRoutingTableFindXClosestSpecificOrder(t *testing.T) {
	testName := "TestRoutingTableFindXClosestSpecificOrder"
	verbose := false
	router := NewRoutingTable([5]uint32{0, 0, 0, 0, 0}, 160, 10)
	for i := 0; i < 10000; i++ {
		router.AddContact(NewRandomContact())
	}
	nodeA := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	nodeB := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 2})
	nodeC := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 3})
	nodeD := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 4})
	nodeE := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 5})
	nodeF := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 6})
	nodeG := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 7})
	nodeH := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 8})
	nodeI := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 9})
	nodeJ := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 10})

	router.AddContact(nodeA)
	router.AddContact(nodeB)
	router.AddContact(nodeC)
	router.AddContact(nodeD)
	router.AddContact(nodeE)
	router.AddContact(nodeF)
	router.AddContact(nodeG)
	router.AddContact(nodeH)
	router.AddContact(nodeI)
	router.AddContact(nodeJ)

	res, err := router.FindXClosest(20, nodeE.ID())
	if err != nil {
		log.Printf("[%s] - failed... %s\n", testName, err.Error())
		t.Fail()
	}
	if verbose {
		log.Printf("[%s]", testName)
		log.Printf("result for findXClosest to %v resulted in", nodeE)
		for _, v := range res {
			log.Printf("contact: %2s %10v\n", "...", v)
		}
	}
	if res[0] != nodeE {
		t.Fail()
		log.Printf("[%s] - nodeE incorrect position", testName)
	}
	if res[1] != nodeD {
		t.Fail()
		log.Printf("[%s] - nodeD incorrect position", testName)
	}
	if res[2] != nodeG {
		t.Fail()
		log.Printf("[%s] - nodeG incorrect position", testName)
	}
	if res[3] != nodeF {
		t.Fail()
		log.Printf("[%s] - nodeF incorrect position", testName)
	}
	if res[4] != nodeA {
		t.Fail()
		log.Printf("[%s] - nodeA incorrect position", testName)
	}
	if res[5] != nodeC {
		t.Fail()
		log.Printf("[%s] - nodeC incorrect position", testName)
	}
	if res[6] != nodeB {
		t.Fail()
		log.Printf("[%s] - nodeB incorrect position", testName)
	}
	if res[7] != nodeI {
		t.Fail()
		log.Printf("[%s] - nodeI incorrect position", testName)
	}
	if res[8] != nodeH {
		t.Fail()
		log.Printf("[%s] - nodeH incorrect position", testName)
	}
	if res[9] != nodeJ {
		t.Fail()
		log.Printf("[%s] - nodeJ incorrect position", testName)
	}

	if t.Failed() {
		for _, v := range res {
			log.Printf("contact: %v", v.ID())
		}
	}
}

func TestRoutingTableDisplay(t *testing.T) {
	testName := "TestRoutingTableDisplay"
	verbose := true
	if verbose {
		rt := NewRoutingTable(RandomID(), KEYSPACE, KBUCKETVOLUME)
		for range 10 {
			rt.AddContact(NewRandomContact())
		}
		log.Printf("[%s]\n%s", testName, rt.Display())
	}
}
