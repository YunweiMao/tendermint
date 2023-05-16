package abcicli

import (
	"fmt"
	"sync"

	tmsync "github.com/YunweiMao/tendermint/libs/sync"
	"github.com/YunweiMao/tendermint/libs/service"

	//RequestInfo is defined in ./abci/types/types.pb.go
	//ResponseInfo is defined in ./abci/types/types.pb.go
	//RequestSetOption is defined in ./abci/types/types.pb.go
	//ResponseSetOption is defined in ./abci/types/types.pb.go
	//RequestDeliverTx is defined in ./abci/types/types.pb.go
	//ResponseDeliverTx is defined in ./abci/types/types.pb.go
	//RequestCheckTx is defined in ./abci/types/types.pb.go
	//ResponseCheckTx is defined in ./abci/types/types.pb.go
	//RequestQuery is defined in ./abci/types/types.pb.go
	//ResponseQuery is defined in ./abci/types/types.pb.go
	//RequestInitChain is defined in ./abci/types/types.pb.go
	//ResponseInitChain is defined in ./abci/types/types.pb.go
	//RequestBeginBlock is defined in ./abci/types/types.pb.go
	//ResponseBeginBlock is defined in ./abci/types/types.pb.go
	//RequestEndBlock is defined in ./abci/types/types.pb.go
	//ResponseEndBlock is defined in ./abci/types/types.pb.go
	//RequestListSnapshots is defined in ./abci/types/types.pb.go
	//ResponseListSnapshots is defined in ./abci/types/types.pb.go
	//RequestOfferSnapshot is defined in ./abci/types/types.pb.go
	//ResponseOfferSnapshot is defined in ./abci/types/types.pb.go
	//RequestLoadSnapshotChunk is defined in ./abci/types/types.pb.go
	//ResponseLoadSnapshotChunk is defined in ./abci/types/types.pb.go
	//RequestApplySnapshotChunk is defined in ./abci/types/types.pb.go
	//ResponseApplySnapshotChunk is defined in ./abci/types/types.pb.go
	//ResponseEcho is defined in ./abci/types/types.pb.go
	//ResponseCommit is defined in ./abci/types/types.pb.go
	//Request is defined in ./abci/types/types.pb.go
	//Response is defined in ./abci/types/types.pb.go
	"github.com/YunweiMao/tendermint/abci/types"
)

const (
	dialRetryIntervalSeconds = 3
	echoRetryIntervalSeconds = 1
)

// Client defines an interface for an ABCI client.
// All `Async` methods return a `ReqRes` object.
// All `Sync` methods return the appropriate protobuf ResponseXxx struct and an error.
// Note these are client errors, eg. ABCI socket connectivity issues.
// Application-related errors are reflected in response via ABCI error codes and logs.
type Client interface {
	service.Service

	SetResponseCallback(Callback)
	Error() error

	FlushAsync() *ReqRes
	EchoAsync(msg string) *ReqRes
	InfoAsync(types.RequestInfo) *ReqRes
	SetOptionAsync(types.RequestSetOption) *ReqRes
	DeliverTxAsync(types.RequestDeliverTx) *ReqRes
	CheckTxAsync(types.RequestCheckTx) *ReqRes
	QueryAsync(types.RequestQuery) *ReqRes
	CommitAsync() *ReqRes
	InitChainAsync(types.RequestInitChain) *ReqRes
	BeginBlockAsync(types.RequestBeginBlock) *ReqRes
	EndBlockAsync(types.RequestEndBlock) *ReqRes
	ListSnapshotsAsync(types.RequestListSnapshots) *ReqRes
	OfferSnapshotAsync(types.RequestOfferSnapshot) *ReqRes
	LoadSnapshotChunkAsync(types.RequestLoadSnapshotChunk) *ReqRes
	ApplySnapshotChunkAsync(types.RequestApplySnapshotChunk) *ReqRes

	FlushSync() error
	EchoSync(msg string) (*types.ResponseEcho, error)
	InfoSync(types.RequestInfo) (*types.ResponseInfo, error)
	SetOptionSync(types.RequestSetOption) (*types.ResponseSetOption, error)
	DeliverTxSync(types.RequestDeliverTx) (*types.ResponseDeliverTx, error)
	CheckTxSync(types.RequestCheckTx) (*types.ResponseCheckTx, error)
	QuerySync(types.RequestQuery) (*types.ResponseQuery, error)
	CommitSync() (*types.ResponseCommit, error)
	InitChainSync(types.RequestInitChain) (*types.ResponseInitChain, error)
	BeginBlockSync(types.RequestBeginBlock) (*types.ResponseBeginBlock, error)
	EndBlockSync(types.RequestEndBlock) (*types.ResponseEndBlock, error)
	ListSnapshotsSync(types.RequestListSnapshots) (*types.ResponseListSnapshots, error)
	OfferSnapshotSync(types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error)
	LoadSnapshotChunkSync(types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunkSync(types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error)
}

//----------------------------------------

// NewClient returns a new ABCI client of the specified transport type.
// It returns an error if the transport is not "socket" or "grpc"
func NewClient(addr, transport string, mustConnect bool) (client Client, err error) {
	switch transport {
	case "socket":
		//NewSocketClient is defined in ./abci/client/socket_client.go
		client = NewSocketClient(addr, mustConnect)
	case "grpc":
		//NewGRPCClient is defined in ./abci/client/grpc_client.go
		client = NewGRPCClient(addr, mustConnect)
	default:
		err = fmt.Errorf("unknown abci transport %s", transport)
	}
	return
}

type Callback func(*types.Request, *types.Response)

type ReqRes struct {
	*types.Request
	*sync.WaitGroup
	*types.Response // Not set atomically, so be sure to use WaitGroup.

	mtx tmsync.Mutex

	// callbackInvoked as a variable to track if the callback was already
	// invoked during the regular execution of the request. This variable
	// allows clients to set the callback simultaneously without potentially
	// invoking the callback twice by accident, once when 'SetCallback' is
	// called and once during the normal request.
	callbackInvoked bool
	cb              func(*types.Response) // A single callback that may be set.
}

func NewReqRes(req *types.Request) *ReqRes {
	return &ReqRes{
		Request:   req,
		WaitGroup: waitGroup1(),
		Response:  nil,

		callbackInvoked: false,
		cb:              nil,
	}
}

// Sets sets the callback. If reqRes is already done, it will call the cb
// immediately. Note, reqRes.cb should not change if reqRes.done and only one
// callback is supported.
func (r *ReqRes) SetCallback(cb func(res *types.Response)) {
	r.mtx.Lock()

	if r.callbackInvoked {
		r.mtx.Unlock()
		cb(r.Response)
		return
	}

	r.cb = cb
	r.mtx.Unlock()
}

// InvokeCallback invokes a thread-safe execution of the configured callback
// if non-nil.
func (r *ReqRes) InvokeCallback() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if r.cb != nil {
		r.cb(r.Response)
	}
	r.callbackInvoked = true
}

// GetCallback returns the configured callback of the ReqRes object which may be
// nil. Note, it is not safe to concurrently call this in cases where it is
// marked done and SetCallback is called before calling GetCallback as that
// will invoke the callback twice and create a potential race condition.
//
// ref: https://github.com/tendermint/tendermint/issues/5439
func (r *ReqRes) GetCallback() func(*types.Response) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return r.cb
}

func waitGroup1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
