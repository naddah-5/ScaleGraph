package scalegraph

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateContact(t *testing.T) {
	var errMsg string = "[TestCreateContact] \n"
	var errMsgDiff string = "[TestCreateContact] \n"
	var expectedIP [4]byte = [4]byte{127, 0, 0, 1}
	var expectedID [5]uint32 = *new([5]uint32)
	newContact := BuildContact(expectedIP, expectedID)
	if newContact.IP() != expectedIP {
		var err string = fmt.Sprintf("IP missmatch: expected - %v, received - %v\n", expectedIP, newContact.IP())
		errMsg = errMsg + err
	}
	if newContact.id != expectedID {
		eID := fmt.Sprintf("%v", expectedID)
		fID := fmt.Sprintf("%v", newContact.id)
		errMsg = errMsg + "ID missmatch: expected - " + eID + " received - " + fID + "\n"
	}
	if errMsg != errMsgDiff {
		log.Println(errMsg)
		t.FailNow()
	}
}
