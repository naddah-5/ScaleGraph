package main

import (
	"errors"
	"strconv"
	"strings"
)

type Contact struct {
	nodeIP  string
	udpPort int
	nodeID  string
}

func NewContact(ip string, udp int, id string) (Contact, error) {
	var newContact Contact
	err := validateIPStructure(ip)
	if err != nil {
		return newContact, err
	}
	newContact = Contact{
		nodeIP:  ip,
		udpPort: udp,
		nodeID:  id,
	}
	return newContact, nil
}

func validateIPStructure(ip string) error {
	var segment []string = strings.Split(ip, ".")
	if len(segment) != 4 {
		return errors.New("invalid ip length")
	}
	for i := 0; i < len(segment); i++ {
		segValue, err := strconv.Atoi(segment[i])
		if err != nil {
			return errors.New("could not parse ip segment: " + err.Error())
		}
		if segValue < 0 || segValue > 255 {
			return errors.New("ip segment out of bounds, valid for 0 <= segment <= 255, found: " + segment[i])
		}
	}
	return nil
}
