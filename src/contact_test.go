package scaleGraph


import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestValidateUDPPortNegative(t *testing.T) {
	var testName string = "TestValidateUDPPortNegative"
	var negativePort int = -80
	err := validateUDPPort(negativePort)
	if err == nil {
		log.Printf("[%s] - %s\n", testName, err.Error())
	}
}

func TestValidateUDPPortLARGE(t *testing.T) {
	var testName string = "TestValidateUDPPortLarge"
	var largePort int = 8080
	err := validateUDPPort(largePort)
	if err == nil {
		log.Printf("[%s] - %s\n", testName, err.Error())
	}
}

func TestCreateContact(t *testing.T) {
	var errMsg string = "[TestCreateContact] \n"
	var errMsgDiff string = "[TestCreateContact] \n"
	var expectedIP [4]byte = [4]byte{127, 0, 0, 1}
	var expectedPort int = 80
	var expectedID [5]uint32 = *new([5]uint32)
	newContact, err := BuildContact(expectedIP, expectedPort, expectedID)
	if err != nil {
		errMsg = errMsg + "could not create new contact " + err.Error() + "\n"
	}
	if newContact.IP() != expectedIP {
		var err string = fmt.Sprintf("IP missmatch: expected - %v, received - %v\n", expectedIP, newContact.IP())
		errMsg = errMsg + err
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
		log.Println(errMsg)
		t.FailNow()
	}
}

func TestNewRandomContact(t *testing.T) {
	var testName string = "TestNewRandomContact"
	var verbose bool = false
	randomContact, err := NewRandomContact()
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.Fail()
	}
	if verbose {
		log.Printf("contact: %+v", randomContact)
	}
}
