package main

import (
	"errors"
	"strconv"
	"strings"
)

type contact struct {
	nodeIP  string
	udpPort int
	nodeID  [5]uint32
}

func (c *contact) IP() string {
	return c.nodeIP
}

func (c *contact) Port() int {
	return c.udpPort
}

func (c *contact) ID() [5]uint32 {
	return c.nodeID
}

func BuildContact(ip string, port int, id [5]uint32) (contact, error) {
	var newContact contact
	var err error = validateContactInfo(ip, port, id)
	if err != nil {
		return newContact, err
	}

	newContact = contact{
		nodeIP:  ip,
		udpPort: port,
		nodeID:  id,
	}
	return newContact, nil
}

// Provides a empty contact instance.
func EmptyContact() contact {
	var empty contact = contact{}
	return empty
}

func validateContactInfo(ip string, port int, id [5]uint32) error {
	var errMsg []error
	err := validateIPStructure(ip)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	err = validateUDPPort(port)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	err = validateNodeID(id)
	if err != nil {
		errMsg = append(errMsg, err)
	}
	if len(errMsg) == 0 {
		return nil
	} else {
		var errString string = ""
		for i := 0; i < len(errMsg); i++ {
			errString += errString + "\n" + errMsg[i].Error()
		}
		err := errors.New(errString)
		return err
	}
}

func validateIPStructure(ip string) error {
	var segment []string = strings.Split(ip, ".")
	var errMsg string = ""
	if len(segment) != 4 {
		errMsg = errMsg + ("invalid ip format, must be in the form of: x.x.x.x received: " + ip + "\n")
	}
	for i := 0; i < len(segment); i++ {
		segValue, err := strconv.Atoi(segment[i])
		if err != nil {
			errMsg = errMsg + ("could not parse ip segment: " + err.Error() + "\n")
		}
		if segValue < 0 || segValue > 255 {
			errMsg = errMsg + "ip segment out of bounds, valid for 0 <= segment <= 255, received: " + segment[i] + " in address " + ip + "\n"
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func validateUDPPort(port int) error {
	if port < 0 || port > 1023 {
		return errors.New("forbidden UDP port submitted, use ports in the range; 0 <= port <= 1023\t received port " + strconv.Itoa(port))
	}
	return nil
}

func validateNodeID(id [5]uint32) error {
	// This will validate id once it is implemented
	// i.e. the signing of it or something like that
	return nil
}
