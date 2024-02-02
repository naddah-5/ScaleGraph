package main

import (
	"testing"
)

func TestNewRoutingTable(t *testing.T) {
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var _ routingTable = NewRoutingTable(homeID)
}
