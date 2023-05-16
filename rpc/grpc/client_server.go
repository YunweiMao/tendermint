package coregrpc

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	//Connect is in libs/net/net.go
	tmnet "github.com/YunweiMao/tendermint/libs/net"
)

// Config is an gRPC server configuration.
type Config struct {
	MaxOpenConnections int
}

// StartGRPCServer starts a new gRPC BroadcastAPIServer using the given
// net.Listener.
// NOTE: This function blocks - you may want to call it in a go-routine.
func StartGRPCServer(ln net.Listener) error {
	grpcServer := grpc.NewServer()
	//RegisterBroadcastAPIServer is defined in rpc/grpc/types.pb.go
	//broadcastAPI is defined in rpc/grpc/api.go 
	RegisterBroadcastAPIServer(grpcServer, &broadcastAPI{})
	return grpcServer.Serve(ln)
}

// StartGRPCClient dials the gRPC server using protoAddr and returns a new
// BroadcastAPIClient.
//BroadcastAPIClient is defined in rpc/grpc/types.pb.go
func StartGRPCClient(protoAddr string) BroadcastAPIClient {
	//nolint:staticcheck // SA1019 Existing use of deprecated but supported dial option.
	conn, err := grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(dialerFunc))
	if err != nil {
		panic(err)
	}
	//NewBroadcastAPIClient is defined in rpc/grpc/types.pb.go
	return NewBroadcastAPIClient(conn)
}

func dialerFunc(ctx context.Context, addr string) (net.Conn, error) {
	return tmnet.Connect(addr)
}
