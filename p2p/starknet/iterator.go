package starknet

import (
	"fmt"

	"github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
)

type iterator struct {
	bcReader blockchain.Reader

	blockNumber uint64
	step        uint64
	limit       uint64
	forward     bool

	err error
}

func newIterator(bcReader blockchain.Reader, blockNumber, limit, step uint64, forward bool) (*iterator, error) {
	if step == 0 {
		return nil, fmt.Errorf("step is zero")
	}

	return &iterator{
		bcReader:    bcReader,
		blockNumber: blockNumber,
		limit:       limit,
		step:        step,
		forward:     forward,
	}, nil
}

func (it *iterator) Err() error {
	return it.err
}

func (it *iterator) Valid() bool {
	if it.limit <= 0 || it.err != nil {
		return false
	}

	return true
}

func (it *iterator) Next() bool {
	if it.forward {
		it.blockNumber += it.step
	} else {
		it.blockNumber -= it.step
	}
	it.limit--

	return it.Valid()
}

func (it *iterator) Block() *core.Block {
	var block *core.Block
	block, it.err = it.bcReader.BlockByNumber(it.blockNumber)

	return block
}
