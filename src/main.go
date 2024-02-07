package main

import (
	"log"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 80
)

func main() {
	log.Println("hello world")
	newTestServer := NewServer([4]byte{127, 0, 0, 1})
	log.Fatal(newTestServer.Start())
}
