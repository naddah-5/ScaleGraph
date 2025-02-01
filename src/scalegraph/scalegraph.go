package scalegraph

import (
	"errors"
	"fmt"
	"sync"
)

type Scalegraph struct {
	sync.RWMutex
	content map[[5]uint32]*Account
}

func NewScaleGraph() *Scalegraph {
	scale := Scalegraph{
		content: make(map[[5]uint32]*Account),
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

	scale.content[id] = NewAccount(id)

	return nil
}

func (scale *Scalegraph) StoredAccountCount() int {
	scale.RLock()
	defer scale.RUnlock()
	return len(scale.content)
}

func (scale *Scalegraph) StoredAccounts() [][5]uint32 {
	scale.RLock()
	defer scale.RUnlock()
	res := make([][5]uint32, 0, len(scale.content))
	for _, accID := range scale.content {
		res = append(res, accID.id)
	}
	return res
}

func (scale *Scalegraph) FindAccount(id [5]uint32) (*Account, error) {
	scale.RLock()
	defer scale.RUnlock()

	account, ok := scale.content[id]
	if !ok {
		return nil, errors.New("account not found")
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
