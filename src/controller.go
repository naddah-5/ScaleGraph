package scalegraph

import (
	"log"
)

// Parses the RPC and dispatch subroutines
func (n *Node) Controller(rpc RPC) {
	switch rpc.CMD {
	case PING:
		n.controlPing(rpc)
	case FIND_NODE:
		n.controlFindNode(rpc)
	case SEND:
		n.controlSend(rpc)
	case STORE:
		n.controlStore(rpc)

	}
}

// Redirection command from the simnet, the redirection target will be stored in the sender.
func (node *Node) controlSend(rpc RPC) {
	rpc.Redirect(rpc.sender.IP())
	rpc.sender = node.contact
	rpc.CMD = rpc.order
}

// Handles the internal logic for a received ping RPC.
func (n *Node) controlPing(rpc RPC) {
	if DEBUG {
		log.Printf("[node] - received %+v, from %+v", rpc.CMD, rpc.sender.id)
	}
	n.AddContact(rpc.sender)
	resp := GenerateResponse(PONG, rpc.id, rpc.sender.ip, n.contact)
	resp.Pong()
	n.network.Send(resp)
}

// Handle the store wallet logic, does NOT overwrite existing wallet or wallet balance.
// i.e. does not overwrite an incorrect chain.
func (n *Node) controlStore(rpc RPC) {
	n.AddContact(rpc.sender)
	err := n.vault.Add(NewWallet(rpc.walletID, rpc.walletBalance))
	if err != nil {
		log.Printf("[INFO] - attempted overwrite of wallet, %v", rpc.walletID)
	}

	resp := GenerateResponse(STORE, rpc.id, rpc.sender.IP(), n.contact)
	resp.StoreResponse()

	n.Send(resp)
}

// Handles the internal logic for find node rpc.
func (n *Node) controlFindNode(rpc RPC) {
	n.AddContact(rpc.sender)
	res, err := n.FindXClosest(REPLICATION, rpc.findTarget)
	if err != nil {
		log.Printf("%+v - find node error: %+v", n.ID(), err)
	}
	resp := GenerateResponse(FIND_NODE, rpc.id, rpc.sender.IP(), n.contact)
	resp.FindNodeResponse(res, rpc.findTarget)
	n.network.Send(resp)
}

func (n *Node) controlShowWallet(rpc RPC) {
	n.AddContact(rpc.sender)
	resp := GenerateResponse(SHOW_WALLET, rpc.id, rpc.sender.IP(), n.contact)

	wallet, err := n.vault.FindWallet(rpc.walletID)
	if err != nil {
		resp.ShowWalletResponse(wallet.walletID, wallet.Balance())
	} else {
		resp.Fail()
	}

}
