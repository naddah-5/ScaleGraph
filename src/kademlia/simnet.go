package kademlia

import (
	"log"
	"sync"
)

// Record and all active IDs and IPs as well as pairwise connections.
type spawned struct {
	id   map[[5]uint32]bool
	ip   map[[4]byte]bool
	node map[[5]uint32][4]byte
	sync.RWMutex
}

type chanTable struct {
	content map[[4]byte]chan RPC
	sync.RWMutex
}

// Simulated network that handles routing between nodes.
// Additionally keeps track of all active nodes.
type Simnet struct {
	spawned
	chanTable
	listener   chan RPC
	serverID   [5]uint32
	serverIP   [4]byte
	masterNode Contact
}

func NewServer() *Simnet {
	s := Simnet{
		spawned: spawned{
			id:   make(map[[5]uint32]bool),
			ip:   make(map[[4]byte]bool),
			node: make(map[[5]uint32][4]byte),
		},
		chanTable: chanTable{
			content: make(map[[4]byte]chan RPC),
		},
		listener: make(chan RPC),
		serverID: [5]uint32{0, 0, 0, 0, 0},
		serverIP: [4]byte{0, 0, 0, 0},
	}

	// Generate master node and attach it to the server.
	master := s.GenerateRandomNode()
	s.masterNode = NewContact(master.IP(), master.ID())

	return &s
}

func (simnet *Simnet) SpawnNode() *Node {
	newNode := simnet.GenerateRandomNode()
	// go newNode.Start()
	return newNode
}

// Generates a new node with random values attaches it to the server and returns a pointer to it.
func (simnet *Simnet) GenerateRandomNode() *Node {
	simnet.spawned.Lock()
	defer simnet.spawned.Unlock()

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

	nodeReceiver := make(chan RPC, 128)
	simnet.chanTable.content[ip] = nodeReceiver
	newNode := NewNode(id, ip, nodeReceiver, simnet.listener, simnet.serverIP, simnet.masterNode)
	return newNode
}

// Initialize listening loop which spawns goroutines.
func (simnet *Simnet) StartServer() {
	for {
		rpc := <-simnet.listener
		go simnet.Route(rpc)
	}
}

// Routes incomming RPC to the correct nodes.
func (simnet *Simnet) Route(rpc RPC) {
	simnet.chanTable.RLock()
	routeChan, ok := simnet.chanTable.content[rpc.Receiver]
	simnet.chanTable.RUnlock()
	if !ok {
		log.Printf("[ERROR] - could not locate node channel for node IP %v", rpc.Receiver)
		return
	}
	routeChan <- rpc
}
