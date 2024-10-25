package src

import (
	"main/src/kademlia"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 3   // alpha
	CONCURRENCY   = 3
	PORT          = 8080
	DEBUG         = true
	POINT_DEBUG   = true
	TIMEOUT       = 10 * time.Second
)

type Node struct {
	kademlia.Contact
	kademlia.Network
	controller chan kademlia.RPC // the channel for internal network, new rpc's are to be sent here for handling
}

