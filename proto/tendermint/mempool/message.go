package mempool

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/YunweiMao/tendermint/p2p"
)

//this file will be used in mempool/v1/reactor.go:198:17
//        					mempool/v0/reactor.go:213:17

var _ p2p.Wrapper = &Txs{}
var _ p2p.Unwrapper = &Message{}

// Wrap implements the p2p Wrapper interface and wraps a mempool message.
func (m *Txs) Wrap() proto.Message {
	mm := &Message{}
	mm.Sum = &Message_Txs{Txs: m}
	return mm
}

// Unwrap implements the p2p Wrapper interface and unwraps a wrapped mempool
// message.
func (m *Message) Unwrap() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Message_Txs:
		return m.GetTxs(), nil

	default:
		return nil, fmt.Errorf("unknown message: %T", msg)
	}
}
