package kademlia

import (
	"fmt"
	"main/src/scalegraph"
)

type cmd int

const (
	NO_CMD cmd = iota
	PING
	PONG
	ENTER
	FIND_NODE
	FOUND_NODES
	INSERT_ACCOUNT
	STORE_ACCOUNT
	STORED_ACCOUNT
	FIND_ACCOUNT
	FOUND_ACCOUNT
	DISPLAY_ACCOUNT
	DISPLAYED_ACCOUNT
	LOCK_ACCOUNT
	LOCKED_ACCOUNT
	UNLOCK_ACCOUNT
	START_TRANSACTION
	PROPOSE_TRANSACTION
	ACCEPT_TRANSACTION
	APPEND_TRANSACTION
)

func (cmd cmd) String() string {
	switch cmd {
	case NO_CMD:
		return "NO_CMD"
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	case ENTER:
		return "ENTER"
	case FIND_NODE:
		return "FIND_NODE"
	case FOUND_NODES:
		return "FOUND_NODES"
	case INSERT_ACCOUNT:
		return "INSERT_ACCOUNT"
	case STORE_ACCOUNT:
		return "STORE_ACCOUNT"
	case STORED_ACCOUNT:
		return "STORED_ACCOUNT"
	case FIND_ACCOUNT:
		return "FIND_ACCOUNT"
	case FOUND_ACCOUNT:
		return "FOUND_ACCOUNT"
	case DISPLAY_ACCOUNT:
		return "DISPLAY_ACCOUNT"
	case DISPLAYED_ACCOUNT:
		return "DISPLAYED_ACCOUNT"
	case LOCK_ACCOUNT:
		return "LOCK_ACCOUNT"
	case LOCKED_ACCOUNT:
		return "LOCKED_ACCOUNT"
	case UNLOCK_ACCOUNT:
		return "UNLOCK_ACCOUNT"
	case START_TRANSACTION:
		return "START_TRANSACTION"
	case PROPOSE_TRANSACTION:
		return "PROPOSE_TRANSACTION"
	case ACCEPT_TRANSACTION:
		return "ACCEPT_TRANSACTION"
	case APPEND_TRANSACTION:
		return "APPEND_TRANSACTION"
	}
	return "unknown cmd"
}

type RPC struct {
	id              [5]uint32
	cmd             cmd
	response        bool
	sender          Contact
	receiver        [4]byte
	findNodeTarget  [5]uint32
	foundNodes      []Contact
	accountID       [5]uint32
	displayString   string
	storeAccSucc    bool
	findAccountSucc bool
	lockChan        chan RPC
	blockID         [5]uint32
	transaction     scalegraph.Transaction
	transactionID   [5]uint32
}

// Generate a fresh send RPC, for a response RPC use GenerateResponse instead.
func GenerateRPC(receiver [4]byte, sender Contact) RPC {
	rpc := RPC{
		id:       RandomID(),
		receiver: receiver,
		response: false,
		sender:   sender,
	}
	return rpc
}

func (rpc *RPC) OverrideID(newID [5]uint32) {
	rpc.id = newID
}

// Generates a fresh response RPC.
func GenerateResponse(id [5]uint32, receiver [4]byte, sender Contact) RPC {
	rpc := RPC{
		id:       id,
		receiver: receiver,
		response: true,
		sender:   sender,
	}
	return rpc
}

// Set a RPC as a ping.
func (rpc *RPC) Ping() {
	rpc.cmd = PING
}

func (rpc *RPC) Pong() {
	rpc.cmd = PONG
}

// Used to get a random existing node in the network from the simulated network.
func (rpc *RPC) Enter() {
	rpc.cmd = ENTER
}

func (rpc *RPC) FindNode(targetNode [5]uint32) {
	rpc.cmd = FIND_NODE
	rpc.findNodeTarget = targetNode
}

func (rpc *RPC) FoundNodes(target [5]uint32, nodes []Contact) {
	rpc.cmd = FOUND_NODES
	rpc.findNodeTarget = target
	rpc.foundNodes = nodes
}

func (rpc *RPC) InsertAccount(accID [5]uint32) {
	rpc.cmd = INSERT_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) StoreAccount(accID [5]uint32) {
	rpc.cmd = STORE_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) StoredAccount(accID [5]uint32, success bool) {
	rpc.cmd = STORED_ACCOUNT
	rpc.accountID = accID
	rpc.storeAccSucc = success
}

func (rpc *RPC) FindAccount(accID [5]uint32) {
	rpc.cmd = FIND_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) FoundAccount(accID [5]uint32, success bool) {
	rpc.cmd = FOUND_ACCOUNT
	rpc.accountID = accID
	rpc.findAccountSucc = success
}

func (rpc *RPC) DisplayAccount(accID [5]uint32) {
	rpc.cmd = DISPLAY_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) DisplayedAccount(accID [5]uint32, displayString string) {
	rpc.cmd = DISPLAYED_ACCOUNT
	rpc.accountID = accID
	rpc.displayString = displayString
}

func (rpc *RPC) LockAccount(accID [5]uint32, lockChan chan RPC) {
	rpc.cmd = LOCK_ACCOUNT
	rpc.accountID = accID
	rpc.lockChan = lockChan
}

func (rpc *RPC) LockedAccount(accID [5]uint32, lockChan chan RPC) {
	rpc.cmd = LOCK_ACCOUNT
	rpc.accountID = accID
	rpc.lockChan = lockChan
}

func (rpc *RPC) UnlockAccount(accID [5]uint32) {
	rpc.cmd = UNLOCK_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) StartTransaction(trx scalegraph.Transaction) {
	rpc.cmd = START_TRANSACTION
	rpc.transaction = trx
}

func (rpc *RPC) ProposeTransaction(trx scalegraph.Transaction) {
	rpc.cmd = PROPOSE_TRANSACTION
	rpc.transaction = trx
}

func (rpc *RPC) AcceptTransaction(trxID [5]uint32) {
	rpc.cmd = ACCEPT_TRANSACTION
	rpc.transactionID = trxID
}

func (rpc *RPC) Display() string {
	rpcString := fmt.Sprintf("id: %v\n", rpc.id)
	rpcString += fmt.Sprintf("CMD: %s\n", rpc.cmd)
	rpcString += fmt.Sprintf("Response: %t\n", rpc.response)
	rpcString += fmt.Sprintf("Sender: %s\n", rpc.sender.Display())
	rpcString += fmt.Sprintf("Receiver: %v\n", rpc.receiver)

	if rpc.cmd == FIND_NODE {
		rpcString += fmt.Sprintf("Find Node Target: %v", rpc.findNodeTarget)
	}
	if rpc.cmd == FOUND_NODES && rpc.response {
		rpcString += "Found Nodes:"
		for _, val := range rpc.foundNodes {
			rpcString += fmt.Sprintf("\n%s", val.Display())
		}
		rpcString += "\n"
	}
	if rpc.cmd == STORE_ACCOUNT {
		rpcString += fmt.Sprintf("store account: %10v\n", rpc.accountID)
	}
	if rpc.cmd == STORED_ACCOUNT && rpc.response {
		rpcString += fmt.Sprintf("stored account: %10v\n", rpc.accountID)
		rpcString += fmt.Sprintf("stored account success: %t", rpc.storeAccSucc)
	}
	rpcString += "\n\n"

	return rpcString
}
