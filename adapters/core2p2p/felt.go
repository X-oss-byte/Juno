package core2p2p

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/p2p/starknet/spec"
	"github.com/NethermindEth/juno/utils"
)

func AdaptHash(f *felt.Felt) *spec.Hash {
	if f == nil {
		return nil
	}

	return &spec.Hash{
		Elements: f.Marshal(),
	}
}

func AdaptSignature(signature []*felt.Felt) *spec.AccountSignature {
	return &spec.AccountSignature{
		Parts: utils.Map(signature, AdaptFelt),
	}
}

func AdaptFelt(f *felt.Felt) *spec.Felt252 {
	if f == nil {
		return nil
	}

	return &spec.Felt252{
		Elements: f.Marshal(),
	}
}

func AdaptFeltSlice(sl []*felt.Felt) []*spec.Felt252 {
	return utils.Map(sl, AdaptFelt)
}

func AdaptFeltToAddress(f *felt.Felt) *spec.Address {
	fBytes := f.Bytes()
	return &spec.Address{
		Elements: fBytes[:],
	}
}
