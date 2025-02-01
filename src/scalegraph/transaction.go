package scalegraph

import "fmt"

// The transaction ID is used as a unique token for the transaction as it is highly
// improbable that two matching ID's are generated randomly.
// Valiators refer to the sending accounts validator nodes while confirmers refer to the receivers
// validator nodes.
type Transaction struct {
	id               [5]uint32
	sendingAccount   [5]uint32
	receivingAccount [5]uint32
	validators       [][5]uint32 // validators for sending account
	confirmers       [][5]uint32 // validators for receiving account
}

func NewTransaction(sender [5]uint32, receriver [5]uint32) *Transaction {
	trx := Transaction{
		id:               RandomID(),
		sendingAccount:   sender,
		receivingAccount: receriver,
	}
	return &trx
}

// Creates a copy of a transaction, this is needed to have copies of the slice's contents
// and not just the pointers to the slices.
func (trx *Transaction) Copy() *Transaction {
	copyValidators := make([][5]uint32, 0, len(trx.validators))
	copyValidators = append(copyValidators, trx.validators...)
	copyConfirmers := make([][5]uint32, 0, len(trx.confirmers))
	copyConfirmers = append(copyConfirmers, trx.confirmers...)
	newTrx := Transaction{
		id:               trx.id,
		sendingAccount:   trx.sendingAccount,
		receivingAccount: trx.receivingAccount,
		validators:       copyValidators,
		confirmers:       copyConfirmers,
	}
	return &newTrx
}

func (trx *Transaction) Display() string {
	disp := ""
	disp += fmt.Sprintf("id: %10v\nsending account: %10v\nreceiving account: %10v", trx.id, trx.sendingAccount, trx.receivingAccount)
	return disp
}
