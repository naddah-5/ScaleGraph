package main

type bucket struct {
	buckets    [KBUCKETS][BUCKETVOLUME]*contact
	bucketSize int
}

func NewBucket() bucket {
	var newBucket bucket = bucket{
		buckets:    [KBUCKETS][BUCKETVOLUME]*contact{},
		bucketSize: KBUCKETS,
	}

	return newBucket
}
