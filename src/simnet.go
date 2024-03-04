package scalegraph

import (
	"log"
	"math/rand"
)

type NodeMap struct {
	ID [5]uint32
	IP [4]byte
}

// Simnet is a simulated network that handles communication between nodes.
// spawnedID and spawnedIP keeps track of id's and ip's that are in use to avoid conflicts
type Simnet struct {
	table     map[[4]byte]chan RPC
	listener  chan RPC
	spawnedID map[[5]uint32]bool
	spawnedIP map[[4]byte]bool
	network   map[[5]uint32][4]byte
	serverID  [5]uint32
	serverIP  [4]byte
}

// Creates a new server
func NewServer() Simnet {
	s := Simnet{
		table:     make(map[[4]byte]chan RPC),
		listener:  make(chan RPC, 100),
		spawnedID: make(map[[5]uint32]bool),
		spawnedIP: make(map[[4]byte]bool),
		network:   make(map[[5]uint32][4]byte),
	}
	rootID := [5]uint32{0, 0, 0, 0, 0}
	rootIP := [4]byte{0, 0, 0, 0}
	s.serverID = rootID
	s.serverIP = rootIP
	s.spawnedID[rootID] = true
	s.spawnedIP[rootIP] = true
	servChan := make(chan RPC, 100)
	s.table[s.serverIP] = servChan

	return s
}

// Spawns a new node and attach it to the server
// Checks for duplicate value conflicts
func (s *Simnet) SpawnNode() {
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
	newNode := NewNode(id, ip, receiver, s.listener, s.serverIP)
	s.table[ip] = receiver
	s.spawnedID[id] = true
	s.spawnedIP[ip] = true
	s.network[id] = ip

	if DEBUG {
		log.Printf("starting node: %+v", newNode.Me.ID())
	}
	go newNode.Start()
}

// Start the server routine, just connects incomming RPC's to the correct channel.
func (s *Simnet) StartServer() {
	for {
		rpc := <-s.listener
		if rpc.Receiver == s.serverIP {
			log.Printf("received a server rpc: %+v", rpc)

		}
		outChan, ok := s.table[rpc.Receiver]
		if !ok {
			log.Printf("received rpc for unknown address, IP: %+v", rpc.Receiver)
		} else {
			outChan <- rpc
		}
	}
}

// Gives all existing nodes with their IP address for the specified network.
// No order is guaranteed.
func (s *Simnet) AllNodes() []NodeMap {
	res := make([]NodeMap, 0)
	for key, value := range s.network {
		nm := NodeMap{key, value}
		res = append(res, nm)
	}
	return res
}

// Redirect pings to the server to a random node in the network.
func (s *Simnet) serverPing(rpc RPC) {
	rng := rand.Int()
	rng %= len(s.spawnedIP)
	rng++
	for value := range s.spawnedIP {
		if 0 == rng {
			rpc.Receiver = value
		}
		rng--
	}
	s.listener <- rpc
}
