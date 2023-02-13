module github.com/aximchain/asc-double-sign-sdk

go 1.13

require (
	github.com/aximchain/go-sdk v0.1.6
	github.com/ethereum/go-ethereum v1.9.13
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.3
	github.com/zondax/ledger-go v0.14.1 // indirect
	golang.org/x/crypto v0.5.0
)

replace (
	github.com/tendermint/tendermint => github.com/bnb-chain/bnc-tendermint v0.32.3-binance.7
	github.com/zondax/hid => github.com/binance-chain/hid v0.9.1-0.20190807012304-e1ffd6f0a3cc
)
