//go:generate protoc --go_out=./ --proto_path=./ --go_opt=Mp2p/proto/requests.proto=./spec --go_opt=Mp2p/proto/transaction.proto=./spec --go_opt=Mp2p/proto/state.proto=./spec --go_opt=Mp2p/proto/snapshot.proto=./spec --go_opt=Mp2p/proto/receipt.proto=./spec --go_opt=Mp2p/proto/mempool.proto=./spec --go_opt=Mp2p/proto/event.proto=./spec --go_opt=Mp2p/proto/block.proto=./spec --go_opt=Mp2p/proto/common.proto=./spec p2p/proto/transaction.proto p2p/proto/state.proto p2p/proto/snapshot.proto p2p/proto/common.proto p2p/proto/block.proto p2p/proto/event.proto p2p/proto/receipt.proto p2p/proto/requests.proto
package starknet

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"github.com/NethermindEth/juno/adapters/core2p2p"
	"github.com/NethermindEth/juno/adapters/p2p2core"
	"github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/p2p/starknet/spec"
	"github.com/NethermindEth/juno/utils"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	bcReader blockchain.Reader
	log      utils.Logger
}

func NewHandler(bcReader blockchain.Reader, log utils.Logger) *Handler {
	return &Handler{
		bcReader: bcReader,
		log:      log,
	}
}

// bufferPool caches unused buffer objects for later reuse.
var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	return buffer
}

func (h *Handler) StreamHandler(stream network.Stream) {
	defer func() {
		if err := stream.Close(); err != nil {
			h.log.Debugw("Error closing stream", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		}
	}()

	buffer := getBuffer()
	defer bufferPool.Put(buffer)

	if _, err := buffer.ReadFrom(stream); err != nil {
		h.log.Debugw("Error reading from stream", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		return
	}

	var req spec.Request
	if err := proto.Unmarshal(buffer.Bytes(), &req); err != nil {
		h.log.Debugw("Error unmarshalling message", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		return
	}

	response, err := h.reqHandler(&req)
	if err != nil {
		h.log.Debugw("Error handling request", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err, "request", req.String())
		return
	}

	for msg, valid := response(); valid; msg, valid = response() {
		if _, err := protodelim.MarshalTo(stream, msg); err != nil { // todo: figure out if we need buffered io here
			h.log.Debugw("Error writing response", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		}
	}
}

func (h *Handler) reqHandler(req *spec.Request) (Stream[proto.Message], error) {
	var singleResponse proto.Message
	var err error
	switch typedReq := req.GetReq().(type) {
	case *spec.Request_GetBlocks:
		return h.HandleGetBlocks(typedReq.GetBlocks)
	case *spec.Request_GetSignatures:
		singleResponse, err = h.HandleGetSignatures(typedReq.GetSignatures)
	case *spec.Request_GetEvents:
		singleResponse, err = h.HandleGetEvents(typedReq.GetEvents)
	case *spec.Request_GetReceipts:
		singleResponse, err = h.HandleGetReceipts(typedReq.GetReceipts)
	case *spec.Request_GetTransactions:
		singleResponse, err = h.HandleGetTransactions(typedReq.GetTransactions)
	default:
		return nil, fmt.Errorf("unhandled request %T", typedReq)
	}

	if err != nil {
		return nil, err
	}
	return StaticStream[proto.Message](singleResponse), nil
}

func (h *Handler) HandleGetBlocks(req *spec.GetBlocks) (Stream[proto.Message], error) {
	// todo: read from bcReader and adapt to p2p type
	count := uint32(0)
	return func() (proto.Message, bool) {
		if count > 3 {
			return nil, false
		}
		count++
		return &spec.BlockHeader{
			State: &spec.Merkle{
				NLeaves: count - 1,
			},
		}, true
	}, nil
}

func (h *Handler) HandleGetSignatures(req *spec.GetSignatures) (*spec.Signatures, error) {
	// todo: read from bcReader and adapt to p2p type
	return &spec.Signatures{
		Id: req.Id,
	}, nil
}

func (h *Handler) HandleGetEvents(req *spec.GetEvents) (*spec.Events, error) {
	block, err := h.blockByID(req.Id)
	if err != nil {
		return nil, err
	}

	var result spec.Events
	for _, receipt := range block.Receipts {
		for _, ev := range receipt.Events {
			event := &spec.Event{
				FromAddress: core2p2p.AdaptFelt(ev.From),
				Keys:        utils.Map(ev.Keys, core2p2p.AdaptFelt),
				Data:        utils.Map(ev.Data, core2p2p.AdaptFelt),
			}

			result.Events = append(result.Events, event)
		}
	}

	return &result, nil
}

func (h *Handler) HandleGetReceipts(req *spec.GetReceipts) (*spec.Receipts, error) {
	// todo: read from bcReader and adapt to p2p type
	magic := 37
	return &spec.Receipts{
		Receipts: make([]*spec.Receipt, magic),
	}, nil
}

func (h *Handler) HandleGetTransactions(req *spec.GetTransactions) (*spec.Transactions, error) {
	block, err := h.blockByID(req.Id)
	if err != nil {
		return nil, err
	}

	resp := new(spec.Transactions)
	for _, transaction := range block.Transactions {
		tx, err := adaptTransaction(transaction)
		if err != nil {
			return nil, err
		}

		resp.Transactions = append(resp.Transactions, tx)
	}

	return resp, nil
}

func adaptTransaction(transaction core.Transaction) (*spec.Transaction, error) {
	var specTx spec.Transaction

	switch tx := transaction.(type) {
	case *core.L1HandlerTransaction:
		specTx.Common = &spec.TransactionCommon{
			Nonce:   adaptFelt(tx.Nonce),
			Version: adaptFelt((*felt.Felt)(tx.Version)),
		}
		specTx.Txn = &spec.Transaction_L1Handler{
			L1Handler: &spec.L1HandlerTransaction{
				Contract:           adaptFeltToAddress(tx.ContractAddress),
				EntryPointSelector: adaptFelt(tx.EntryPointSelector),
				Calldata:           utils.Map(tx.CallData, adaptFelt),
			},
		}
		return &specTx, nil
	case *core.InvokeTransaction:
		specTx.Common = &spec.TransactionCommon{
			Nonce:   adaptFelt(tx.Nonce),
			Version: adaptFelt((*felt.Felt)(tx.Version)),
		}
		specTx.Txn = &spec.Transaction_L2Transaction{
			L2Transaction: &spec.L2Transaction{
				Common: &spec.L2TransactionCommon{
					Sender:    adaptFeltToAddress(tx.SenderAddress),
					Signature: adaptTxSignature(tx),
					MaxFee:    adaptFelt(tx.MaxFee),
				},
				Txn: &spec.L2Transaction_Invoke{
					Invoke: &spec.InvokeTransaction{
						Calldata: utils.Map(tx.CallData, adaptFelt),
					},
				},
			},
		}
		return &specTx, nil
	case *core.DeclareTransaction:
		specTx.Common = &spec.TransactionCommon{
			Nonce:   adaptFelt(tx.Nonce),
			Version: adaptFelt((*felt.Felt)(tx.Version)),
		}
		specTx.Txn = &spec.Transaction_L2Transaction{
			L2Transaction: &spec.L2Transaction{
				Common: &spec.L2TransactionCommon{
					Sender:    adaptFeltToAddress(tx.SenderAddress),
					Signature: adaptTxSignature(tx),
					MaxFee:    adaptFelt(tx.MaxFee),
				},
				Txn: &spec.L2Transaction_Declare{
					Declare: &spec.DeclareTransaction{
						// ClassHash: adaptFelt(tx.ClassHash),
						// CompiledHash: adaptFelt(tx.CompiledClassHash),
					},
				},
			},
		}
		return &specTx, nil
	default:
		return nil, fmt.Errorf("unsupported tx type %T", tx)
	}
}

func adaptTxSignature(tx core.Transaction) *spec.Signature {
	return &spec.Signature{
		Parts: utils.Map(tx.Signature(), adaptFelt),
	}
}

func adaptFeltToAddress(f *felt.Felt) *spec.Address {
	fBytes := f.Bytes()
	return &spec.Address{
		Elements: fBytes[:],
	}
}

func adaptFelt(f *felt.Felt) *spec.Felt252 {
	fBytes := f.Bytes()
	return &spec.Felt252{
		Elements: fBytes[:],
	}
}

func (h *Handler) blockByID(id *spec.BlockID) (*core.Block, error) {
	switch {
	case id == nil:
		return nil, errors.New("block id is nil")
	case id.Hash != nil:
		hash := p2p2core.AdaptHash(id.Hash)
		return h.bcReader.BlockByHash(hash)
	default:
		return h.bcReader.BlockByNumber(id.Height)
	}
}
