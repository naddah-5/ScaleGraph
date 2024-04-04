package scalegraph

import (
	"log"
	"testing"
)

// Sanity test for the add and remove to a vault.
func TestBaseVault(t *testing.T) {
	testName := "TestBaseVault"
	verbose := false
	testVault := NewVault()
	testWallet := NewWallet(GenerateID(), 0)
	testVault.Add(testWallet)

	res, err := testVault.FindWallet(testWallet.walletID)
	if err != nil {
		log.Printf("[%s] - could not retrieve test wallet\n", testName)
		t.FailNow()
	}

	if verbose {
		log.Printf("found wallet %+v\n", res)
	}

	testVault.Remove(testWallet.walletID)
	res, err = testVault.FindWallet(testWallet.walletID)
	if err != nil {
		log.Printf("[%s] - found removed wallet\n", testName)
		t.FailNow()
	}

	if verbose {
		log.Println("removed wallet")
	}
}
