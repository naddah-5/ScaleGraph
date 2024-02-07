package scaleGraph

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 80
)

type Node struct {
	Alpha    int
	K        int
	KeySpace int
	ID       [5]uint32
	routingTable
}
