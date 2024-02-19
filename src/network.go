package scalegraph

type network struct {
	listener chan RPC
	sender   chan RPC
}

func NewNetwork(ln chan RPC, sn chan RPC) network {
	newNetwork := network{
		listener: ln,
		sender: sn,
	}
	return newNetwork
}

func (n *network) Send(RPC) {}

func (n *network) Listen() {}
