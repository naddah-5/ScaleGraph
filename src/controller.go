package scalegraph

import (
	"log"
)

// Parses the RPC and dispatch subroutines
func (node *Node) Controller(rpc RPC) {
	switch rpc.CMD {
	case PING:
		node.controlPing(rpc)
	case FIND_NODE:
		node.controlFindNode(rpc)
	case SEND:
		node.controlSend(rpc)
	case STORE:
		node.controlStore(rpc)

	}
}

// Redirection command from the simnet, the redirection target will be stored in the sender.
func (node *Node) controlSend(rpc RPC) {
	rpc.Redirect(rpc.sender.IP())
	rpc.sender = node.contact
	rpc.CMD = rpc.order
}

// Handles the internal logic for a received ping RPC.
func (node *Node) controlPing(rpc RPC) {
	node.AddContact(rpc.sender)
	resp := GenerateResponse(PONG, rpc.id, rpc.sender.ip, node.contact)
	resp.Pong()
	node.network.Send(resp)
}

// Handle the store wallet logic, does NOT overwrite existing wallet or wallet balance.
// i.e. does not overwrite an incorrect chain.
func (node *Node) controlStore(rpc RPC) {
	node.AddContact(rpc.sender)
	wallet := NewWallet(rpc.walletID, rpc.walletBalance)
	err := node.vault.Add(wallet)
	if err != nil {
		log.Printf("[INFO] - attempted overwrite of wallet, %v\n", rpc.walletID)
	} else {
		log.Printf("[INFO] - Stored wallet: %v\n", wallet.walletID)
	}

	resp := GenerateResponse(STORE, rpc.id, rpc.sender.IP(), node.contact)
	resp.StoreResponse()

	node.Send(resp)
}

// Handles the internal logic for find node rpc.
func (node *Node) controlFindNode(rpc RPC) {
	node.AddContact(rpc.sender)
	res, err := node.FindXClosest(REPLICATION, rpc.findTarget)
	if err != nil {
		log.Printf("%+v - find node error: %+v", node.ID(), err)
	}
	resp := GenerateResponse(FIND_NODE, rpc.id, rpc.sender.IP(), node.contact)
	resp.FindNodeResponse(res, rpc.findTarget)
	node.network.Send(resp)
}

func (node *Node) controlShowWallet(rpc RPC) {
	node.AddContact(rpc.sender)
	resp := GenerateResponse(SHOW_WALLET, rpc.id, rpc.sender.IP(), node.contact)

	wallet, err := node.vault.FindWallet(rpc.walletID)
	if err != nil {
		resp.ShowWalletResponse(wallet.walletID, wallet.Balance())
	} else {
		resp.Fail()
	}

}

