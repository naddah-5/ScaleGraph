package kademlia

import (
	"math/rand"
)

type Contact struct {
	ip   [4]byte
	id   [5]uint32
	port int
}

func (contact *Contact) IP() [4]byte {
	return contact.ip
}

func (contact *Contact) ID() [5]uint32 {
	return contact.id
}

func NewContact(ip [4]byte, id [5]uint32) Contact {
	contact := Contact{
		ip: ip,
		id: id,
	}
	return contact
}

func NewRandomContact() Contact {
	var ip [4]byte
	var id [5]uint32
	for i := 0; i < 4; i++ {
		seg, _ := RandU32(0, 256)
		ip[i] = byte(seg)
	}
	for i := 0; i < 5; i++ {
		id[i] = rand.Uint32()
	}
	return NewContact(ip, id)
}
