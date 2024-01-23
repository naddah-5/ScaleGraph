package main

import "fmt"

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 80
)

func main() {
	fmt.Println("hello world")
}
