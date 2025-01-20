package scalegraph

import (
	"errors"
	"fmt"
	"sync"
)


type Scalegraph struct {
	sync.RWMutex
	content map[[5]uint32]bool
}

func NewVault() *Scalegraph {
	scale := Scalegraph{
		content: make(map[[5]uint32]bool),
	}
	return &scale
}

func (scale *Scalegraph) AddAccount(id [5]uint32) error {
	scale.Lock()
	defer scale.Unlock()

	_, ok := scale.content[id]
	if ok {
		return errors.New("account already exists")
	}

	scale.content[id] = true

	return nil
}

func (scale *Scalegraph) FindAccount(id [5]uint32) (bool, error) {
	scale.RLock()
	defer scale.RUnlock()

	account, ok := scale.content[id]
	if !ok {
		return false, errors.New("account not found")
	}
	return account, nil
}

func (scale *Scalegraph) RemoveAccount(id [5]uint32) {
	scale.Lock()
	defer scale.Unlock()

	delete(scale.content, id)
}

func (scale *Scalegraph) Display() string {
	scale.Lock()
	defer scale.Unlock()

	view := ""
	for id, acc := range scale.content {
		view += fmt.Sprintf("Account id: %v\ncontent: %v\n", id, acc)
	}

	return view
}
