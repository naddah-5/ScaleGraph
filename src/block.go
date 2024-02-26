package scalegraph

type block struct {
	height           int
	blockHash        []byte
	prevSenderHash   []byte
	prevReceiverHash []byte
	balance          int
	transaction
	concensus [][5]uint32
}
