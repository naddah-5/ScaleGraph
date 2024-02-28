package scalegraph

type block struct {
	heightSender     int
	heightReceiver   int
	blockHash        []byte // the block hash is a hash of the transaction only
	prevSenderHash   []byte
	prevReceiverHash []byte
	balance          int
	transaction
	concensus [][5]uint32
}
