package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestValidateIPNegative(t *testing.T) {
	var testName string = "TestValidateIPNegative"
	var invalidIP string = "-127.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
		t.FailNow()
	}
}

func TestValidateIPLarge(t *testing.T) {
	var testName string = "TestValidateIPLarge"
	var invalidIP string = "1270.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
		t.FailNow()
	}
}

func TestValidateIPText(t *testing.T) {
	var testName string = "TestValidateIPText"
	var invalidIP string = "127.zero.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
		t.FailNow()
	}
}

func TestValidateIPTextManyFaults(t *testing.T) {
	var testName string = "TestValidateIPTextManyFaults"
	var invalidIP string = "1270.zero.0.-1.10"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
		t.FailNow()
	}
}

func TestValidateUDPPortNegative(t *testing.T) {
	var testName string = "TestValidateUDPPortNegative"
	var negativePort int = -80
	err := validateUDPPort(negativePort)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
	}
}

func TestValidateUDPPortLARGE(t *testing.T) {
	var testName string = "TestValidateUDPPortLarge"
	var largePort int = 8080
	err := validateUDPPort(largePort)
	if err == nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
	}
}

func TestCreateContact(t *testing.T) {
	var errMsg string = "[TestCreateContact] \n"
	var errMsgDiff string = "[TestCreateContact] \n"
	var expectedIP string = "127.0.0.1"
	var expectedPort int = 80
	var expectedID [5]uint32 = *new([5]uint32)
	newContact, err := BuildContact(expectedIP, expectedPort, expectedID)
	if err != nil {
		errMsg = errMsg + "could not create new contact " + err.Error() + "\n"
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

func TestNewRandomContact(t *testing.T) {
	var testName string = "TestNewRandomContact"
	var verbose bool = false
	randomContact, err := NewRandomContact()
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.Fail()
	}
	if verbose {
		fmt.Printf("contact: %+v", randomContact)
	}
}
