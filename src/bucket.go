package main

type Bucket struct {
	kBucket    []contact
	bucketSize int
}

type contact struct {
	nodeIP  string
	udpPort int
	nodeID  string
}

func NewBucket(k int) Bucket {
	var newBucket Bucket = Bucket{
		kBucket: make([]contact, k),
		bucketSize: k,
	}

	return newBucket
}
