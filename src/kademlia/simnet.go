package kademlia

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
)

// Record and all active IDs and IPs as well as pairwise connections.
type spawned struct {
	id    map[[5]uint32]bool
	ip    map[[4]byte]bool
	nodes []Contact
	sync.RWMutex
}

type chanTable struct {
	content map[[4]byte]chan RPC
	sync.RWMutex
}

// Simulated network that handles routing between nodes.
// Additionally keeps track of all active nodes.
type Simnet struct {
	chanTable
	spawned
	listener          chan RPC
	serverID          [5]uint32
	serverIP          [4]byte
	masterNode        *Node
	masterNodeContact Contact
	dropPercent       float32
	debug             bool
}

func NewServer(debugMode bool, dropPercent float32) *Simnet {
	s := Simnet{
		chanTable: chanTable{
			content: make(map[[4]byte]chan RPC),
		},
		spawned: spawned{
			id:    make(map[[5]uint32]bool),
			ip:    make(map[[4]byte]bool),
			nodes: make([]Contact, 0),
		},
		listener:    make(chan RPC, 64),
		serverID:    [5]uint32{0, 0, 0, 0, 0},
		serverIP:    [4]byte{0, 0, 0, 0},
		dropPercent: dropPercent,
		debug:       debugMode,
	}

	// Generate master node and attach it to the server.
	s.masterNode = s.GenerateRandomNode()
	s.masterNodeContact = NewContact(s.masterNode.ip, s.masterNode.id)
	// looks stupid but the master node should know that it is in fact the master node.
	s.masterNode.masterNode = s.masterNodeContact

	return &s
}

// Roll the RNG to determine if the rpc should be dropped.
func (simnet *Simnet) DropRoll() bool {
	if simnet.dropPercent == 0.0 {
		return false
	}
	roll := rand.Float32() < simnet.dropPercent
	if roll {
		return true
	}
	return false
}

func (simnet *Simnet) MasterNode() *Node {
	return simnet.masterNode
}

func (simnet *Simnet) SpawnNode(done chan [5]uint32) *Node {
	newNode := simnet.GenerateRandomNode()
	go newNode.Start(done)
	return newNode
}

// Removes node from simnet records and sends a shutdown signal to it.
func (simnet *Simnet) ShudownNode(node *Node) {
	simnet.chanTable.Lock()
	simnet.spawned.Lock()
	defer simnet.chanTable.Unlock()
	defer simnet.spawned.Unlock()

	delete(simnet.chanTable.content, node.IP())
	delete(simnet.spawned.ip, node.IP())
	delete(simnet.spawned.id, node.ID())
	i := slices.Index(simnet.spawned.nodes, node.Contact)
	if i != -1 {
		simnet.spawned.nodes = slices.Delete(simnet.spawned.nodes, i, i+1)
	}
	node.shutdown <- struct{}{}
}

// Generates a new node with random values attaches it to the server and returns a pointer to it.
func (simnet *Simnet) GenerateRandomNode() *Node {
	simnet.chanTable.Lock()
	simnet.spawned.Lock()
	defer simnet.spawned.Unlock()
	defer simnet.chanTable.Unlock()

	id := RandomID()
	_, ok := simnet.spawned.id[id]
	// if the generated id is already taken, generate new ones until a free one is found.
	for ok {
		id = RandomID()
		_, ok = simnet.spawned.id[id]
	}
	simnet.spawned.id[id] = true

	ip := RandomIP()
	_, ok = simnet.spawned.ip[ip]
	// if the generated ip is already taken, generate new ones until a free one is found.
	for ok {
		ip = RandomIP()
		_, ok = simnet.spawned.ip[ip]
	}
	simnet.spawned.ip[ip] = true

	node := NewContact(ip, id)
	simnet.spawned.nodes = append(simnet.spawned.nodes, node)

	nodeReceiver := make(chan RPC, 128)
	simnet.chanTable.content[ip] = nodeReceiver
	newNode := NewNode(id, ip, nodeReceiver, simnet.listener, simnet.serverIP, simnet.masterNodeContact, false)
	return newNode
}

// Returns contact information for a random node in the network.
func (simnet *Simnet) randomNode() Contact {
	simnet.spawned.RLock()
	defer simnet.spawned.RUnlock()
	index, _ := RandU32(0, uint32(len(simnet.nodes)))
	return simnet.nodes[index]
}

// Initialize listening loop which spawns goroutines.
func (simnet *Simnet) StartServer() {
	// Master node should not be part of the main wait group.
	go simnet.masterNode.Start(make(chan [5]uint32, 64))
	for {
		rpc := <-simnet.listener
		if simnet.debug {
			log.Printf("[DEBUG] - simnet queue: %d", len(simnet.listener))
		}
		go simnet.Route(rpc)
	}
}

func (simnet *Simnet) ListKnownIPChannels() string {
	simnet.chanTable.RLock()
	defer simnet.chanTable.RUnlock()
	keys := make([][4]byte, 0, len(simnet.chanTable.content))
	for k := range simnet.chanTable.content {
		keys = append(keys, k)
	}
	keyString := fmt.Sprint("known IP channels:")
	for _, val := range keys {
		keyString += fmt.Sprintf("\n%v", val)
	}
	return keyString
}

// Routes incomming RPC to the correct nodes.
func (simnet *Simnet) Route(rpc RPC) {
	simnet.chanTable.RLock()
	defer simnet.chanTable.RUnlock()

	routeChan, ok := simnet.chanTable.content[rpc.receiver]
	if !ok {
		log.Printf("[ERROR] - could not locate node channel for node IP %v RPC %s", rpc.receiver, rpc.Display())
		return
	}

	if rpc.cmd == ENTER {
		simnet.spawned.RLock()
		defer simnet.spawned.RUnlock()
		node := simnet.randomNode()
		for node.ID() == [5]uint32{0, 0, 0, 0, 0} {
			node = simnet.masterNodeContact
		}
		rpc.foundNodes = append(rpc.foundNodes, node)
		rpc.response = true
	}

	if simnet.DropRoll() {
		if simnet.debug {
			log.Printf("Dropping RPC: %v\n", rpc.id)
		}
		return
	}
	routeChan <- rpc
	return
}
