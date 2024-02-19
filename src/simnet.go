package scalegraph

// Simnet is a simulated network that handles communication between nodes.
// spawnedID and spawnedIP keeps track of id's and ip's that are in use to avoid conflicts
type Simnet struct {
	table     map[[4]byte]chan RPC
	listener  chan RPC
	spawnedID map[[5]uint32]bool
	spawnedIP map[[4]byte]bool
	serverIP  [4]byte
}

// Creates a new server
func NewServer() Simnet {
	s := Simnet{
		table:     make(map[[4]byte]chan RPC),
		listener:  make(chan RPC),
		spawnedID: make(map[[5]uint32]bool),
		spawnedIP: make(map[[4]byte]bool),
	}
	root := [4]byte{0, 0, 0, 0}
	s.serverIP = root
	s.spawnedIP[root] = true

	return s
}

// Spawns a new node and attach it to the server
// Checks for duplicate value conflicts
func (s *Simnet) SpawnNode() {
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
	receiver := make(chan RPC)
	newNode := NewNode(id, ip, receiver, s.listener)
	s.table[ip] = receiver
	s.spawnedID[id] = true
	s.spawnedIP[ip] = true

	go newNode.Start(s.serverIP)
}
