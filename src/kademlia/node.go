package kademlia

import (
	"fmt"
	"main/src/scalegraph"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 20  // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	CONCURRENCY   = REPLICATION
	PORT          = 8080
	DEBUG         = true
	POINT_DEBUG   = true
	TIMEOUT       = 500 * time.Millisecond
)

type Node struct {
	Contact
	Network
	RoutingTable
	scalegraph scalegraph.Scalegraph
	controller chan RPC // the channel for internal network, new rpc's are to be sent here for handling
	shutdown   chan struct{}
	debug      bool
}

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC, serverIP [4]byte, masterNode Contact, debug bool) *Node {
	controller := make(chan RPC)
	net := NewNetwork(id, listener, sender, controller, serverIP, masterNode, false)
	me := NewContact(ip, id)
	router := NewRoutingTable(me, KEYSPACE, KBUCKETVOLUME)
	return &Node{
		Contact:      me,
		Network:      *net,
		RoutingTable: *router,
		scalegraph:   *scalegraph.NewScaleGraph(),
		shutdown:     make(chan struct{}),
		debug:        debug,
	}
}

// Starts up the node, joining the network via the "Enter", and "Find node" protocols.
func (node *Node) Start(done chan [5]uint32) {
	go node.Network.Listen(node)
	if node.Contact.IP() == node.masterNode.IP() {
		return
	} else {
		node.Enter()
		done <- node.ID()
	}
}

// Wrapper for sending a rpc and also adding the responding contact.
func (node *Node) Send(rpc RPC) (RPC, error) {
	res, err := node.Network.Send(rpc)
	if err != nil {
		// If the contact fails to respond and exists in the routing table, drop it.
		con, ipErr := node.FindByIP(rpc.receiver)
		if ipErr == nil {
			node.RemoveContact(con)
		}
		return res, err
	} else {
		node.AddContact(res.sender)
		return res, nil
	}
}

func (node *Node) Debug(mode bool) {
	node.debug = mode
	node.Network.Debug(mode)
}

func (node *Node) AddAccount(id [5]uint32) error {
	err := node.scalegraph.AddAccount(id)
	if err != nil {
		return err
	}
	return nil
}

func (node *Node) ClearDeadContacts() {
	contacts := node.RoutingTable.AllContacts()
	for _, con := range contacts {
		go node.Ping(con.IP())
	}
	time.Sleep(TIMEOUT)
}

func (node *Node) Display() string {
	res := fmt.Sprintf("node ID routing table state: node IP %v\tnode ID: %v", node.IP(), node.ID())
	res += node.RoutingTable.Display()
	return res
}

