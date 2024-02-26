package scalegraph

type transaction struct {
	Sender    [5]uint32
	Recipient [5]uint32
	Ammount   int
	Signature []byte
}
