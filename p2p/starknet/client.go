package starknet

import (
	"context"

	"github.com/NethermindEth/juno/p2p/starknet/spec"
	"github.com/NethermindEth/juno/utils"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/proto"
)

type NewStreamFunc func(ctx context.Context, pids ...protocol.ID) (network.Stream, error)

type Client struct {
	newStream  NewStreamFunc
	protocolID protocol.ID
	log        utils.Logger
}

func NewClient(newStream NewStreamFunc, protocolID protocol.ID, log utils.Logger) *Client {
	return &Client{
		newStream:  newStream,
		protocolID: protocolID,
		log:        log,
	}
}

func (c *Client) sendAndCloseWrite(stream network.Stream, req proto.Message) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	if _, err = stream.Write(reqBytes); err != nil {
		return err
	}
	return stream.CloseWrite()
}

func (c *Client) receiveInto(stream network.Stream, res proto.Message) error {
	return protodelim.UnmarshalFrom(&byteReader{stream}, res)
}

func (c *Client) sendAndReceiveInto(ctx context.Context, req, res proto.Message) error {
	stream, err := c.newStream(ctx, c.protocolID)
	if err != nil {
		return err
	}
	defer stream.Close() // todo: dont ignore close errors

	if err = c.sendAndCloseWrite(stream, req); err != nil {
		return err
	}

	return c.receiveInto(stream, res)
}

func (c *Client) GetBlockBodies(ctx context.Context, it *spec.Iteration) (Stream[*spec.BlockHeader], error) {
	wrappedReq := spec.BlockBodiesRequest{
		Iteration: it,
	}

	stream, err := c.newStream(ctx, c.protocolID)
	if err != nil {
		return nil, err
	}
	if err := c.sendAndCloseWrite(stream, &wrappedReq); err != nil {
		return nil, err
	}

	return func() (*spec.BlockHeader, bool) {
		var res spec.BlockHeader
		if err := c.receiveInto(stream, &res); err != nil {
			stream.Close() // todo: dont ignore close errors
			return nil, false
		}
		return &res, true
	}, nil
}

// func (c *Client) GetSignatures(ctx context.Context, it *spec.Iteration) (*spec.Signatures, error) {
//	wrappedReq := spec.Request{
//		Req: &spec.Request_GetSignatures{
//			GetSignatures: req,
//		},
//	}
//
//	var res spec.Signatures
//	if err := c.sendAndReceiveInto(ctx, &wrappedReq, &res); err != nil {
//		return nil, err
//	}
//	return &res, nil
// }

func (c *Client) GetEvents(ctx context.Context, it *spec.Iteration) (*spec.EventsResponse, error) {
	wrappedReq := spec.EventsRequest{
		Iteration: it,
	}

	var res spec.EventsResponse
	if err := c.sendAndReceiveInto(ctx, &wrappedReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetReceipts(ctx context.Context, it *spec.Iteration) (*spec.ReceiptsResponse, error) {
	wrappedReq := spec.ReceiptsRequest{
		Iteration: it,
	}

	var res spec.ReceiptsResponse
	if err := c.sendAndReceiveInto(ctx, &wrappedReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetTransactions(ctx context.Context, it *spec.Iteration) (*spec.TransactionsResponse, error) {
	wrappedReq := spec.TransactionsRequest{
		Iteration: it,
	}

	var res spec.TransactionsResponse
	if err := c.sendAndReceiveInto(ctx, &wrappedReq, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
