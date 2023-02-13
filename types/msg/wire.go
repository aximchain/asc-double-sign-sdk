package msg

import (
	"github.com/aximchain/go-sdk/types/msg"
	"github.com/aximchain/go-sdk/types/tx"
	"github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(MsgAscSubmitEvidence{}, "cosmos-sdk/MsgAscSubmitEvidence", nil)
}

func init() {
	RegisterCodec(msg.MsgCdc)
	RegisterCodec(tx.Cdc)
}
