package main

import (
	"log"
	scaleGraph "scalegraph/src"
)



func main() {
	log.Println("hello world")
	_ = scaleGraph.NewServer([4]byte{127, 0, 0, 1})
}
