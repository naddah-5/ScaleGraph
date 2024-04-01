package scalegraph

import (
	"log"
	"sync"
)

// Used to keep track of existing node id's and ip's and their connection.
type spawned struct {
	spawnLock sync.RWMutex
	spawnedID map[[5]uint32]bool
	spawnedIP map[[4]byte]bool
	network   map[[5]uint32][4]byte
}

type table struct {
	tableLock sync.RWMutex
	content   map[[4]byte]chan RPC
}

// Simnet is a simulated network that handles communication between nodes.
// spawnedID and spawnedIP keeps track of id's and ip's that are in use to avoid conflicts
type Simnet struct {
	listener chan RPC
	spawned
	table
	serverID   [5]uint32
	serverIP   [4]byte
	masterNode [4]byte
}

// Creates a new server
func NewServer() *Simnet {
	s := Simnet{
		listener: make(chan RPC, 1000),
		spawned: spawned{
			spawnLock: sync.RWMutex{},
			spawnedID: make(map[[5]uint32]bool),
			spawnedIP: make(map[[4]byte]bool),
			network:   make(map[[5]uint32][4]byte),
		},
		table: table{
			tableLock: sync.RWMutex{},
			content:   make(map[[4]byte]chan RPC),
		},
	}
	rootID := [5]uint32{0, 0, 0, 0, 0}
	rootIP := [4]byte{0, 0, 0, 0}
	s.serverID = rootID
	s.serverIP = rootIP
	s.spawnedID[rootID] = true
	s.spawnedIP[rootIP] = true
	servChan := make(chan RPC, 100)
	s.table.content[s.serverIP] = servChan

	master := s.SpawnNode()
	s.masterNode = master.IP()

	return &s
}

// Spawns a new node and attach it to the server
// Checks for duplicate value conflicts
func (s *Simnet) SpawnNode() *Node {
	newNode := s.generateRandomNode()
	if DEBUG {
		log.Printf("generated id: %+v, ip: %+v", newNode.ID(), newNode.IP())
	}
	go newNode.Start()
	return newNode
}

// Generates a new node with random ID, and IP in the simulation.
func (s *Simnet) generateRandomNode() *Node {
	s.spawnLock.Lock()
	defer s.spawnLock.Unlock()
	if DEBUG {
		log.Println("spawning node")
	}
	var id [5]uint32
	for {
		id = GenerateID()
		_, ok := s.spawnedID[id]
		if !ok {
			break
		}
	}
	var ip [4]byte
	for {
		ip = GenerateIP()
		_, ok := s.spawnedIP[ip]
		if !ok {
			break
		}
	}

	receiver := make(chan RPC, 100)
	newNode := NewNode(id, ip, receiver, s.listener, s.serverIP, s.masterNode)
	s.table.content[ip] = receiver
	s.spawnedID[id] = true
	s.spawnedIP[ip] = true
	s.network[id] = ip

	return &newNode
}

// Attaches a generated node to the simulation.
// Does not handle id or ip conflicts.
func (s *Simnet) AttachThisNode(node *Node) {
	s.spawnLock.Lock()
	defer s.spawnLock.Unlock()
	s.table.content[node.IP()] = node.listener
	s.spawnedID[node.ID()] = true
	s.spawnedIP[node.IP()] = true
	s.network[node.ID()] = node.IP()
}

// Start the server routine, just connects incomming RPC's to the correct channel.
func (s *Simnet) StartServer() {
	if DEBUG {
		log.Printf("starting server with id: %+v, ip: %+v", s.serverID, s.serverIP)
	}
	for {
		rpc := <-s.listener
		go s.understand(rpc)
	}
}

// Used to launch a go routine for handling incoming traffic.
func (s *Simnet) understand(rpc RPC) {
	if rpc.receiver == s.serverIP {
		if DEBUG {
			if rpc.CMD == PONG {
				log.Printf("[warning] - server received a PONG")
			}
		}
		if DEBUG {
			log.Printf("[server] - received a server rpc: %+v", rpc)
		}
		s.serverPing(rpc)
	}
	s.tableLock.RLock()
	outChan, ok := s.table.content[rpc.receiver]
	if !ok {
		log.Printf("[server] - received rpc for unknown address, IP: %+v, sender: %+v", rpc.receiver, rpc.Sender.id)
	} else {
		outChan <- rpc
	}
	s.tableLock.RUnlock()
	return
}

// Gives all existing nodes with their IP address for the specified network.
// No order is guaranteed.
func (s *Simnet) AllNodes() []contact {
	s.spawnLock.RLock()
	defer s.spawnLock.RUnlock()
	res := make([]contact, 0)
	for key, value := range s.network {
		nm := contact{value, key}
		res = append(res, nm)
	}
	return res
}

// Redirect pings to the server to a random node in the network.
func (s *Simnet) serverPing(rpc RPC) {
	rpc.receiver = s.masterNode
	if DEBUG {
		log.Printf("[server] - redirecting %+v\n\tid: %+v \n\tsender: %+v \n\treceiver: %+v", rpc, rpc.ID, rpc.Sender, rpc.receiver)
	}
	s.listener <- rpc
}
