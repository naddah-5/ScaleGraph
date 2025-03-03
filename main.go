package main

import (
	"fmt"
	"main/src/kademlia"
)

func main() {
	testSize := 50
	batchSize := 20
	for i := testSize; i > 35; i -= 5 {
		testIteration(i, batchSize)
	}
}

func testIteration(testSize int, batchSize int) {
	testData := make([]int, kademlia.REPLICATION)
	totalMisses := 0
	fmt.Printf("================== TEST START ==================\n")
	fmt.Printf("\t\t\t\t\tn = %d\n", testSize)
	for i := range batchSize {
		miss, res := kademlia.IntegrationNodeLookupDataGathering(testSize)
		fmt.Printf("test #%d:\n", i+1)
		fmt.Printf("%d incomplete validator groups\ntest result: %v\n", miss, res)
		totalMisses += miss
		for i, missing := range res {
			testData[i] += missing
		}
	}
	fmt.Println("================== TEST COMPLETE ==================")
	fmt.Println("")
	fmt.Printf("%d incomplete validator groups across all tests for n = %d\ndistribution: %v\n", totalMisses, testSize, testData)
	fmt.Println("")
	fmt.Println("")

}
