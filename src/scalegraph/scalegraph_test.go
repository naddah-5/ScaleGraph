package scalegraph

import (
	"log"
	"testing"
)

func TestScalegraphDisplay(t *testing.T) {
	testName := "TestScalegraphDisplay"
	verbose := false
	sg := NewScaleGraph()
	for range 5 {
		sg.AddAccount(RandomID())
	}
	view := sg.Display()
	if verbose {
		log.Printf("[%s]\n%s", testName, view)
	}
}
