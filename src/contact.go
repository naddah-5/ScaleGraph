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

func (c *Contact) IP() string {
	return c.nodeIP
}

func (c *Contact) Port() int {
	return c.udpPort
}

func (c *Contact) ID() string {
	return c.nodeID
}

func NewContact(ip string, port int, id string) (Contact, error) {
	var newContact Contact
	var conErr error = validateContactInfo(ip, port, id)
	if conErr != nil {
		return newContact, conErr
	}

	newContact = Contact{
		nodeIP:  ip,
		udpPort: port,
		nodeID:  id,
	}
	return newContact, nil
}

func validateContactInfo(ip string, port int, id string) error {
	var errMsg []error
	ipErr := validateIPStructure(ip)
	if ipErr != nil {
		errMsg = append(errMsg, ipErr)
	}
	udpErr := validateUDPPort(port)
	if udpErr != nil {
		errMsg = append(errMsg, udpErr)
	}
	nodeErr := validateNodeID(id)
	if nodeErr != nil {
		errMsg = append(errMsg, nodeErr)
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
	if len(segment) != 4 {
		return errors.New("invalid ip format, must be in the form of: x.x.x.x\t received: " + ip)
	}
	for i := 0; i < len(segment); i++ {
		segValue, err := strconv.Atoi(segment[i])
		if err != nil {
			return errors.New("could not parse ip segment: " + err.Error())
		}
		if segValue < 0 || segValue > 255 {
			return errors.New("ip segment out of bounds, valid for 0 <= segment <= 255, received: " + segment[i] + " in address " + ip)
		}
	}
	return nil
}

func validateUDPPort(port int) error {
	if port < 0 || port > 1023 {
		return errors.New("forbidden UDP port submitted, use ports in the range; 0 <= port <= 1023\t received port " + strconv.Itoa(port))
	}
	return nil
}

func validateNodeID(id string) error {
	// This will validate id once it is implemented
	return nil
}
