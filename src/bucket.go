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

func (b *bucket) AddContact(ip string, port int, id [5]uint32) error {
	_, genErr := NewContact(ip, port, id)
	if genErr != nil {
		return genErr
	}
	// relDistance, calcErr := _
	return nil
}
