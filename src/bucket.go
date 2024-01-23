package main

type bucket struct {
	content [KBUCKETVOLUME]contact
}

func NewBucket() *bucket {
	var newBucket bucket = bucket{
		content: [KBUCKETVOLUME]contact{},
	}

	return &newBucket
}

func (b *bucket) AddContact(ip string, port int, id [5]uint32) error {
	_, genErr := NewContact(ip, port, id)
	if genErr != nil {
		return genErr
	}
	// relDistance, := _
	return nil
}
