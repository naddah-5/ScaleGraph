package scalegraph

import "sync"

type Account struct {
	sync.RWMutex
	blockChain []Block
}


