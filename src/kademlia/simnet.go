package kademlia

import (
	"fmt"
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
	listener          chan RPC
	serverID          [5]uint32
	serverIP          [4]byte
	masterNode        *Node
	masterNodeContact Contact
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
	s.masterNode = s.GenerateRandomNode()
	s.masterNodeContact = NewContact(s.masterNode.ip, s.masterNode.id)
	// looks stupid but the master node should know that it is in fact the master node.
	s.masterNode.masterNode = s.masterNodeContact

	return &s
}

func (simnet *Simnet) SpawnNode() *Node {
	newNode := simnet.GenerateRandomNode()
	go newNode.Start()
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
	newNode := NewNode(id, ip, nodeReceiver, simnet.listener, simnet.serverIP, simnet.masterNodeContact)
	return newNode
}

// Initialize listening loop which spawns goroutines.
func (simnet *Simnet) StartServer() {
	go simnet.masterNode.Start()
	for {
		rpc := <-simnet.listener
		go simnet.Route(rpc)
	}
}

func (simnet *Simnet) DebugKnownIPChannels() string {
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
	routeChan, ok := simnet.chanTable.content[rpc.Receiver]
	simnet.chanTable.RUnlock()
	if !ok {
		log.Printf("[ERROR] - could not locate node channel for node IP %v", rpc.Receiver)
		return
	}
	routeChan <- rpc
}
