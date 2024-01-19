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
		fmt.Println("[TestValidateIPNegative] \n" + err.Error())
		t.FailNow()
	}
}

func TestValidateIPLarge(t *testing.T) {
	var invalidIP string = "1270.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPLarge] \n" + err.Error())
		t.FailNow()
	}
}

func TestValidateIPText(t *testing.T) {
	var invalidIP string = "127.zero.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPText] \n" + err.Error())
		t.FailNow()
	}
}

func TestValidateIPTextManyFaults(t *testing.T) {
	var invalidIP string = "1270.zero.0.-1.10"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println("[TestValidateIPTextManyFaults] \n" + err.Error())
		t.FailNow()
	}
}

func TestValidateUDPPortNegative(t *testing.T) {
	var negativePort int = -80
	err := validateUDPPort(negativePort)
	if err == nil {
		fmt.Println("[TestValidateUDPPortNegative] \n" + err.Error())
	}
}

func TestValidateUDPPortLARGE(t *testing.T) {
	var largePort int = 8080
	err := validateUDPPort(largePort)
	if err == nil {
		fmt.Println("[TestValidateUDPPortLarge] \n" + err.Error())
	}
}

func TestCreateContact(t *testing.T) {
	var errMsg string = "[TestCreateContact] \n"
	var errMsgDiff string = "[TestCreateContact] \n"
	var expectedIP string = "127.0.0.1"
	var expectedPort int = 80
	var expectedID [5]uint32 = *new([5]uint32)
	newContact, conErr := NewContact(expectedIP, expectedPort, expectedID)
	if conErr != nil {
		errMsg = errMsg + "could not create new contact " + conErr.Error() + "\n"
	}
	if newContact.IP() != expectedIP {
		errMsg = errMsg + "IP missmatch: expected - " + expectedIP + " received - " + newContact.nodeIP + "\n"
	}
	if newContact.udpPort != expectedPort {
		errMsg = errMsg + "Port missmatch: expected - " + strconv.Itoa(expectedPort) + " received - " + strconv.Itoa(newContact.Port()) + "\n"
	}
	if newContact.nodeID != expectedID {
		eID := fmt.Sprintf("%v", expectedID)
		fID := fmt.Sprintf("%v", newContact.nodeID)
		errMsg = errMsg + "ID missmatch: expected - " + eID + " received - " + fID + "\n"
	}
	if errMsg != errMsgDiff {
		fmt.Println(errMsg)
		t.FailNow()
	}
}
