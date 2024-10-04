package scalegraph

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

func (node *Node) Ping(ip [4]byte) error {
	ping := GenerateRPC(PING, node.contact, ip)
	resp, err := node.network.Send(ping)
	if err != nil {
		return errors.New(fmt.Sprintf("no ping response from IP: %+v", ip))
	}
	node.AddContact(resp.sender)
	return nil
}

// High level find node RPC.
func (node *Node) FindNode(target [5]uint32) []contact {
	closeIP, _ := node.routingTable.FindXClosest(REPLICATION, target)

	res := list.New()
	respChan := make(chan []contact, CONCURRENCY)

	for n := closeIP.Front(); n != nil; n = n.Next() {
		closeNode := n.Value.(contact)
		rpc := GenerateRPC(FIND_NODE, node.contact, closeNode.ip)
		rpc.FindNode(closeNode.ID())
		go node.alphaFindNode(rpc, respChan)
	}

	for i := 0; i < min(CONCURRENCY, closeIP.Len()); i++ {
		foundNodes := <-respChan
		if DEBUG {
			log.Printf("[info] - received find node alpha response: %+v", foundNodes)
		}
		for _, val := range foundNodes {
			res.PushBack(val)
		}
	}

	// TODO: Replace list sort with a slice sort
	SortByDistance(res, target)
	var resSlice []contact = make([]contact, 0, REPLICATION)
	for c := res.Front(); c != nil; c = c.Next() {
		if len(resSlice) >= REPLICATION {
			break
		}
		resSlice = append(resSlice, c.Value.(contact))
	}
	finalRes := node.deepSearch(resSlice, target)
	return finalRes
}

func (node *Node) alphaFindNode(rpc RPC, respChan chan []contact) {
	resp, err := node.network.Send(rpc)
	if err != nil {
		log.Println("[error] - find node alpha: no response")
		respChan <- make([]contact, 0)
		return
	}
	if DEBUG {
		log.Println("[info] - sending find node alpha response")
	}
	respChan <- resp.kNodes
	go func(rpc RPC) {
		for _, con := range rpc.kNodes {
			go node.Ping(con.IP())
		}
	}(resp)
	return

}

// Helper function to the find node RPC, handles the reccursion.
func (node *Node) deepSearch(prevContactList []contact, target [5]uint32) []contact {
	var respChan chan []contact = make(chan []contact, CONCURRENCY)
	for i := 0; i < len(prevContactList); i++ {
		rpc := GenerateRPC(FIND_NODE, node.contact, prevContactList[i].IP())
		rpc.FindNode(target)
		go node.alphaFindNode(rpc, respChan)
	}

	res := list.New()
	for i := 0; i < min(CONCURRENCY, len(prevContactList)); i++ {
		foundNodes := <-respChan
		for _, val := range foundNodes {
			res.PushBack(val)
		}
	}
	SortByDistance(res, target)
	contactList := make([]contact, 0, REPLICATION)
	i := 0
	for c := res.Front(); c != nil; c = c.Next() {
		if i >= REPLICATION {
			break
		}
		contactList = append(contactList, c.Value.(contact))
		i++
	}

	if len(prevContactList) > 0 && len(contactList) > 0 {
		if prevContactList[0] == contactList[0] {
			return prevContactList
		}
	} else if len(prevContactList) == 0 && len(contactList) == 0 {
		return contactList
	}
	return node.deepSearch(contactList, target)
}

func (node *Node) StoreWallet(wallet *wallet) error {
	valGroup := node.FindNode(wallet.walletID)
	var wg sync.WaitGroup

	for _, con := range valGroup {
		wg.Add(1)
		go func(walletID [5]uint32, con contact, wg *sync.WaitGroup) {
			walletRPC := GenerateRPC(STORE, node.contact, con.IP())
			walletRPC.Store(walletID, 0)
			resp, err := node.network.Send(walletRPC)
			if err != nil {
				log.Printf("ERROR: store wallet with id %v at node %v timed out with error - %s", walletID, con.ID(), err.Error())
			}
			if resp.acknowledge == false {
				errMsg := "FATAL: store wallet RPC response was incorrectly formated, acknowledge not true"
				data := fmt.Sprintf("Sender %v, Store node %v, Wallet ID %v", node.contact.ID(), con.ID(), walletID)
				log.Printf("%s\n%s", errMsg, data)
				os.Exit(1)
			}
			wg.Done()
		}(wallet.walletID, con, &wg)

	}

	wg.Wait()
	return nil
}

func (node *Node) ShowWallet(walletID [5]uint32) {
	valGroup := node.FindNode(walletID)

	for _, holder := range valGroup {
		rpc := GenerateRPC(SHOW_WALLET, node.contact, holder.IP())
		rpc.ShowWallet(walletID)
		_, err := node.network.Send(rpc)
		if err == nil {
			// add data extraction
		}
	}
	// no wallet error
}
