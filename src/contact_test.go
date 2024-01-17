package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestValidateIPNegative(t *testing.T) {
	var invalidIP string = "-127.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPNegative] - " + err.Error())
		t.FailNow()
	}
}

func TestValidateIPLarge(t *testing.T) {
	var invalidIP string = "1270.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPLarge] - " + err.Error())
		t.FailNow()
	}
}

func TestValidateIPText(t *testing.T) {
	var invalidIP string = "127.zero.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPText] - " + err.Error())
		t.FailNow()
	}
}

func TestValidateUDPPortNegative(t *testing.T) {
	var negativePort int = -80
	err := validateUDPPort(negativePort)
	if err == nil {
		fmt.Println("[TestValidateUDPPortNegative] - " + err.Error())
	}
}

func TestValidateUDPPortLARGE(t *testing.T) {
	var largePort int = 8080
	err := validateUDPPort(largePort)
	if err == nil {
		fmt.Println("[TestValidateUDPPortLarge] - " + err.Error())
	}
}

func TestCreateContact(t *testing.T) {
	var expectedIP string = "127.0.0.1"
	var expectedPort int = 80
	var expectedID string = "node ID"
	newContact, conErr := NewContact(expectedIP, expectedPort, expectedID)
	if conErr != nil {
		fmt.Println("[TestCreateContact] - could not create new contact " + conErr.Error())
		t.Fail()
	}
	if newContact.IP() != expectedIP {
		fmt.Println("[TestCreateContact] - IP missmatch: expected - " + expectedIP + " received - " + newContact.nodeIP)
		t.Fail()
	}
	if newContact.udpPort != expectedPort {
		fmt.Println("[TestCreateContact] - port missmatch: expected - " + strconv.Itoa(expectedPort) + " received - " + strconv.Itoa(newContact.Port()))
	}
	if newContact.nodeID != expectedID {
		fmt.Println("[TestCreateContact] - ID missmatch: expected - " + expectedID + " received - " + newContact.nodeID)
	}
}
