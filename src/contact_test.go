package main

import (
	"fmt"
	"testing"
)

func TestValidateIPNegative(t *testing.T) {
	var invalidIP string = "-127.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println(err)
		t.FailNow()
	}
}


func TestValidateIPLarge(t *testing.T) {
	var invalidIP string = "1270.0.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestValidateIPText(t *testing.T) {
	var invalidIP string = "127.zero.0.1"
	err := validateIPStructure(invalidIP)
	if err == nil {
		fmt.Println(err)
		t.FailNow()
	}
}
