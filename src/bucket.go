package main

type Bucket struct {
	buckets    [KBUCKETS][BUCKETVOLUME]Contact
	bucketSize int
}

func NewBucket() Bucket {
	var newBucket Bucket = Bucket{
		buckets:    [KBUCKETS][BUCKETVOLUME]Contact{},
		bucketSize: KBUCKETS,
	}

	return newBucket
}
