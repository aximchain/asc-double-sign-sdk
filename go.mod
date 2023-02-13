module github.com/aximchain/asc-double-sign-sdk

go 1.13

require (
	github.com/aximchain/go-sdk v0.1.7-fc067ad
	github.com/ethereum/go-ethereum v1.9.13
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563 // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.3
	github.com/zondax/ledger-go v0.14.1 // indirect
	golang.org/x/crypto v0.5.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/tendermint/tendermint => github.com/bnb-chain/bnc-tendermint v0.32.3-binance.7
	github.com/zondax/hid => github.com/binance-chain/hid v0.9.1-0.20190807012304-e1ffd6f0a3cc
)
