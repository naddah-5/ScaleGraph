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

func (node *Node) alphaFindNode(rpc RPC, responseChannel chan []contact) {
	response, err := node.network.Send(rpc)
	if err != nil {
		log.Println("[error] - find node alpha: no response")
		responseChannel <- make([]contact, 0)
		return
	}
	responseChannel <- response.kNodes
	go func(rpc RPC) {
		for _, con := range rpc.kNodes {
			go node.Ping(con.IP())
		}
	}(response)
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
	SortListByDistance(res, target)
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
	s := fmt.Sprintln("intention to store wallet in the following nodes")
	for _, con := range valGroup {
		s += fmt.Sprintf("node ID: %v\n", con.ID())
	}
	log.Println(s)

	for _, con := range valGroup {
		wg.Add(1)
		go func(walletID [5]uint32, con contact, wg *sync.WaitGroup) {
			rpc := GenerateRPC(STORE, node.contact, con.IP())
			rpc.Store(walletID, 0)
			resp, err := node.network.Send(rpc)
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

func (node *Node) ShowWallet(walletID [5]uint32) (string, error) {
	valGroup := node.FindNode(walletID)
	var res string = ""

	for _, holder := range valGroup {
		rpc := GenerateRPC(SHOW_WALLET, node.contact, holder.IP())
		rpc.ShowWallet(walletID)
		resp, err := node.network.Send(rpc)
		if err == nil {
			res = fmt.Sprintf("wallet: %v\nbalance: %d", resp.walletID, resp.walletBalance)
			return res, nil
		}
	}
	return res, errors.New(fmt.Sprintf("[ERROR] - Could not find wallet %v", walletID))
}

func (node *Node) FindNode(target [5]uint32) []contact {
	// find the closest internal nodes
	closeNodes, _ := node.routingTable.FindXClosest(REPLICATION, target)

	// send find node RPC to those nodes
	responses := make(chan []contact, REPLICATION)
	for n := closeNodes.Front(); n != nil; n = n.Next() {
		aNode := n.Value.(contact)
		rpc := GenerateRPC(FIND_NODE, node.contact, aNode.IP())
		rpc.FindNode(target)
		go node.alphaFindNode(rpc, responses)
	}

	// extract the returned contacts
	foundNodes := make([]contact, 0, REPLICATION*REPLICATION)
	for i := 0; i < min(CONCURRENCY, closeNodes.Len()); i++ {
		foundNodes = append(foundNodes, <-responses...)
	}

	// sort the result with a bias towards larger nodes
	SortSliceByDistance(&foundNodes, target)

	// start recursion
	recIndex := min(REPLICATION, len(foundNodes))
	finalRes := node.searchProtocol(foundNodes[:recIndex], target)

	return finalRes
}

func (node *Node) searchProtocol(prevContactList []contact, target [5]uint32) []contact {
	respChan := make(chan []contact)

	for _, v := range prevContactList {
		rpc := GenerateRPC(FIND_NODE, node.contact, v.IP())
		rpc.FindNode(target)
		go node.alphaFindNode(rpc, respChan)
	}

	respList := make([]contact, REPLICATION*REPLICATION)
	for i := 0; i < min(len(prevContactList), CONCURRENCY); i++ {
		respList = append(respList, <-respChan...)
	}

	SortSliceByDistance(&respList, target)

	if RelativeDistance(prevContactList[0].ID(), target) == RelativeDistance(respList[0].ID(), target) {
		return respList
	} else {
		return node.searchProtocol(respList, target)
	}
}
