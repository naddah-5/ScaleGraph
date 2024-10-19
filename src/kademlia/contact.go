package kademlia

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

