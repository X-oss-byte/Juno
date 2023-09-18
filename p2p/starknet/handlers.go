//go:generate protoc --go_out=./ --proto_path=./ --go_opt=Mp2p/proto/transaction.proto=./spec --go_opt=Mp2p/proto/state.proto=./spec --go_opt=Mp2p/proto/snapshot.proto=./spec --go_opt=Mp2p/proto/receipt.proto=./spec --go_opt=Mp2p/proto/mempool.proto=./spec --go_opt=Mp2p/proto/event.proto=./spec --go_opt=Mp2p/proto/block.proto=./spec --go_opt=Mp2p/proto/common.proto=./spec p2p/proto/transaction.proto p2p/proto/state.proto p2p/proto/snapshot.proto p2p/proto/common.proto p2p/proto/block.proto p2p/proto/event.proto p2p/proto/receipt.proto
package starknet

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"github.com/NethermindEth/juno/adapters/core2p2p"
	"github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
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

	var req proto.Message
	if err := proto.Unmarshal(buffer.Bytes(), req); err != nil {
		h.log.Debugw("Error unmarshalling message", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		return
	}

	response, err := h.reqHandler(req)
	if err != nil {
		h.log.Debugw("Error handling request", "peer", stream.ID(),
			"protocol", stream.Protocol(), "err", err, "request", "") // todo req.String() ?
		return
	}

	for msg, valid := response(); valid; msg, valid = response() {
		if _, err := protodelim.MarshalTo(stream, msg); err != nil { // todo: figure out if we need buffered io here
			h.log.Debugw("Error writing response", "peer", stream.ID(), "protocol", stream.Protocol(), "err", err)
		}
	}
}

func (h *Handler) reqHandler(req proto.Message) (Stream[proto.Message], error) {
	var singleResponse proto.Message
	var err error
	switch typedReq := req.(type) {
	case *spec.BlockBodiesRequest:
		return h.HandleGetBlocks(typedReq.Iteration)
	// case *spec.:
	//	singleResponse, err = h.HandleGetSignatures(typedReq.GetSignatures)
	case *spec.EventsRequest:
		singleResponse, err = h.HandleGetEvents(typedReq.Iteration)
	case *spec.ReceiptsRequest:
		singleResponse, err = h.HandleGetReceipts(typedReq.Iteration)
	case *spec.TransactionsRequest:
		return h.HandleGetTransactions(typedReq.Iteration)
	default:
		return nil, fmt.Errorf("unhandled request %T", typedReq)
	}

	if err != nil {
		return nil, err
	}
	return StaticStream[proto.Message](singleResponse), nil
}

func (h *Handler) HandleGetBlocks(i *spec.Iteration) (Stream[proto.Message], error) {
	it, err := h.blockIterator(i)
	if err != nil {
		return nil, err
	}

	return func() (proto.Message, bool) {
		if !it.Valid() {
			return nil, false
		}

		block := it.Block()
		if it.Err() != nil {
			// todo log an error
			return nil, false
		}
		it.Next()

		// todo fill the spec later
		return &spec.BlockBodiesResponse{
			BlockNumber: block.Number,
		}, false
	}, nil
}

func (h *Handler) HandleGetSignatures(req *spec.Iteration) (*spec.Signatures, error) {
	// todo: read from bcReader and adapt to p2p type
	return &spec.Signatures{
		// Id: req.Id,
	}, nil
}

func (h *Handler) HandleGetEvents(req *spec.Iteration) (*spec.EventsResponse, error) {
	block, err := h.blockByIteration(req)
	if err != nil {
		return nil, err
	}

	var events []*spec.Event
	for _, receipt := range block.Receipts {
		for _, ev := range receipt.Events {
			event := &spec.Event{
				FromAddress: core2p2p.AdaptFelt(ev.From),
				Keys:        core2p2p.AdaptFeltSlice(ev.Keys),
				Data:        core2p2p.AdaptFeltSlice(ev.Data),
			}

			events = append(events, event)
		}
	}

	return &spec.EventsResponse{
		BlockNumber: block.Number,
		BlockHash:   core2p2p.AdaptHash(block.Hash),
		Responses: &spec.EventsResponse_Events_{
			Events: &spec.EventsResponse_Events{
				Items: events,
			},
		},
	}, nil
}

func (h *Handler) HandleGetReceipts(req *spec.Iteration) (*spec.ReceiptsResponse, error) {
	// todo: read from bcReader and adapt to p2p type
	magic := uint64(37) //nolint:gomnd
	return &spec.ReceiptsResponse{
		BlockNumber: magic,
		BlockHash:   nil,
		Responses:   nil,
	}, nil
}

func (h *Handler) HandleGetTransactions(i *spec.Iteration) (Stream[proto.Message], error) {
	it, err := h.blockIterator(i)
	if err != nil {
		return nil, err
	}

	return func() (proto.Message, bool) {
		if !it.Valid() {
			return nil, false
		}

		block := it.Block()
		if it.Err() != nil {
			h.log.Errorw("Iterator failure", "err", it.Err())
			return nil, false
		}
		it.Next()

		return &spec.TransactionsResponse{
			BlockNumber: block.Number,
			BlockHash:   core2p2p.AdaptHash(block.Hash),
			Responses: &spec.TransactionsResponse_Transactions{
				Transactions: &spec.Transactions{
					Items: utils.Map(block.Transactions, core2p2p.AdaptTransaction),
				},
			},
		}, false
	}, nil
}

func (h *Handler) blockIterator(it *spec.Iteration) (*iterator, error) {
	forward := it.Direction == spec.Iteration_Forward
	return newIterator(h.bcReader, it.StartBlock, it.Limit, it.Step, forward)
}

func (h *Handler) blockByIteration(it *spec.Iteration) (*core.Block, error) {
	switch {
	case it == nil:
		return nil, errors.New("iteration is nil")
	default:
		return h.bcReader.BlockByNumber(it.StartBlock)
	}
}
