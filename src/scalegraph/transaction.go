package scalegraph

type Transaction struct {
	id               [5]uint32
	sendingAccount   [5]uint32
	receivingAccount [5]uint32
	validators       [][5]uint32
}
