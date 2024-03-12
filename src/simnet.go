package scalegraph

import (
	"log"
	"sync"
)

type NodeMap struct {
	ID [5]uint32
	IP [4]byte
}

type spawned struct {
	lock      sync.RWMutex
	spawnedID map[[5]uint32]bool
	spawnedIP map[[4]byte]bool
	network   map[[5]uint32][4]byte
}

// Simnet is a simulated network that handles communication between nodes.
// spawnedID and spawnedIP keeps track of id's and ip's that are in use to avoid conflicts
type Simnet struct {
	listener chan RPC
	table    map[[4]byte]chan RPC
	spawned
	serverID   [5]uint32
	serverIP   [4]byte
	masterNode [4]byte
}

// Creates a new server
func NewServer() *Simnet {
	s := Simnet{
		table:    make(map[[4]byte]chan RPC),
		listener: make(chan RPC, 1000),
		spawned: spawned{
			lock:      sync.RWMutex{},
			spawnedID: make(map[[5]uint32]bool),
			spawnedIP: make(map[[4]byte]bool),
			network:   make(map[[5]uint32][4]byte),
		},
	}
	s.lock.Lock()
	rootID := [5]uint32{0, 0, 0, 0, 0}
	rootIP := [4]byte{0, 0, 0, 0}
	s.serverID = rootID
	s.serverIP = rootIP
	s.spawnedID[rootID] = true
	s.spawnedIP[rootIP] = true
	servChan := make(chan RPC, 100)
	s.table[s.serverIP] = servChan
	s.lock.Unlock()
	master := s.SpawnNode()
	s.masterNode = master

	return &s
}

// Spawns a new node and attach it to the server
// Checks for duplicate value conflicts
func (s *Simnet) SpawnNode() [4]byte {
	s.lock.Lock()
	defer s.lock.Unlock()
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
	s.table[ip] = receiver
	s.spawnedID[id] = true
	s.spawnedIP[ip] = true
	s.network[id] = ip

	if DEBUG {
		log.Printf("starting node: %+v with ip:%+v", newNode.ID(), newNode.IP())
		log.Printf("generated id: %+v, ip: %+v", id, ip)
	}
	go newNode.Start()
	return newNode.IP()
}

// Start the server routine, just connects incomming RPC's to the correct channel.
func (s *Simnet) StartServer() {
	if DEBUG {
		log.Printf("\tstarting server with id: %+v, ip: %+v", s.serverID, s.serverIP)
	}
	for {
		rpc := <-s.listener
		go s.understand(rpc)
	}
}

func (s *Simnet) understand(rpc RPC) {
	if rpc.receiver == s.serverIP {
		if DEBUG {
			if rpc.CMD == PONG {
				log.Printf("\t---WARNING: server received a PONG---")
			}
		}
		if DEBUG {
			log.Printf("[server] - received a server rpc: %+v", rpc)
		}
		s.serverPing(rpc)
	}
	s.lock.RLock()
	outChan, ok := s.table[rpc.receiver]
	s.lock.RUnlock()
	if !ok {
		log.Printf("[server] - received rpc for unknown address, IP: %+v, sender: %+v", rpc.receiver, rpc.Sender.id)
	} else {
		outChan <- rpc
	}
}

// Gives all existing nodes with their IP address for the specified network.
// No order is guaranteed.
func (s *Simnet) AllNodes() []NodeMap {
	s.lock.Lock()
	defer s.lock.Unlock()
	res := make([]NodeMap, 0)
	for key, value := range s.network {
		nm := NodeMap{key, value}
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
