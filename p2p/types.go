package p2p

import (
	"github.com/gogo/protobuf/proto"

	//ChannelDescriptor and ConnectionStatus are defined in p2p/conn/connection.go
	"github.com/YunweiMao/tendermint/p2p/conn"

	//PexAddrs and PexRequest are defined in /proto/tendermint/p2p/pex.proto/pex.pb.go
	//pex.proto depends on /proto/tendermint/p2p/types.proto
	tmp2p "github.com/YunweiMao/tendermint/proto/tendermint/p2p"
)

type ChannelDescriptor = conn.ChannelDescriptor
type ConnectionStatus = conn.ConnectionStatus

// Envelope contains a message with sender routing info.
type Envelope struct {
	Src       Peer          // sender (empty if outbound)
	Message   proto.Message // message payload
	ChannelID byte
}

// Unwrapper is a Protobuf message that can contain a variety of inner messages
// (e.g. via oneof fields). If a Channel's message type implements Unwrapper, the
// p2p layer will automatically unwrap inbound messages so that reactors do not have to do this themselves.
type Unwrapper interface {
	proto.Message

	// Unwrap will unwrap the inner message contained in this message.
	Unwrap() (proto.Message, error)
}

// Wrapper is a companion type to Unwrapper. It is a Protobuf message that can contain a variety of inner messages. The p2p layer will automatically wrap outbound messages so that the reactors do not have to do it themselves.
type Wrapper interface {
	proto.Message

	// Wrap will take the underlying message and wrap it in its wrapper type.
	Wrap() proto.Message
}

var (
	_ Wrapper = &tmp2p.PexRequest{}
	_ Wrapper = &tmp2p.PexAddrs{}
)
