package scalegraph


// The transaction ID is used as a unique token for the transaction as it is highly
// improbable that two matching ID's are generated randomly.
// Valiators refer to the sending accounts validator nodes while confirmers refer to the receivers
// validator nodes.
type Transaction struct {
	id               [5]uint32
	sendingAccount   [5]uint32
	receivingAccount [5]uint32
	validators       [][5]uint32
	confirmers       [][5]uint32
}

func NewTransaction(sender [5]uint32, receriver [5]uint32) *Transaction {
	trx := Transaction{
		id: RandomID(),
		sendingAccount: sender,
		receivingAccount: receriver,
	}
	return &trx
}


