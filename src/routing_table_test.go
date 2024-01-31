package main

import (
	"fmt"
	"testing"
)

func TestNewRoutingTable(t *testing.T) {
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testRT routingTable = NewRoutingTable(homeID)
	fmt.Printf("%+v", testRT)
}
