package main

import (
	"fmt"
	"testing"
)

func TestReadStructure(t *testing.T) {
	content, err := ReadEnvironment()
	if err != nil {
		fmt.Println("could not retrieve file", err)
		t.FailNow()
	}
	fmt.Println(content)
}
