package client

import (
	asc "github.com/aximchain/asc-double-sign-sdk/types/asc"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func EthHeaderToAscHeader(ethHeader *ethtypes.Header) *asc.Header {
	if ethHeader == nil {
		return nil
	}
	return &asc.Header{
		ParentHash:  ethHeader.ParentHash,
		UncleHash:   ethHeader.UncleHash,
		Coinbase:    ethHeader.Coinbase,
		Root:        ethHeader.Root,
		TxHash:      ethHeader.TxHash,
		ReceiptHash: ethHeader.ReceiptHash,
		Bloom:       ethHeader.Bloom,
		Difficulty:  ethHeader.Difficulty.Int64(),
		Number:      ethHeader.Number.Int64(),
		GasLimit:    ethHeader.GasLimit,
		GasUsed:     ethHeader.GasUsed,
		Time:        ethHeader.Time,
		Extra:       ethHeader.Extra,
		MixDigest:   ethHeader.MixDigest,
		Nonce:       ethHeader.Nonce,
	}
}