package client

import (
	asc "github.com/aximchain/asc-double-sign-sdk/types/asc"
	"github.com/aximchain/asc-double-sign-sdk/types/msg"
	"github.com/aximchain/go-sdk/client/rpc"
	"github.com/aximchain/go-sdk/common/types"
	"github.com/aximchain/go-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func ASCSubmitEvidence(c *rpc.HTTP, submitter types.AccAddress, headers []*asc.Header,
	syncType rpc.SyncType, options ...tx.Option) (*coretypes.ResultBroadcastTx, error) {
	m := msg.NewMsgAscSubmitEvidence(submitter, headers)
	return c.Broadcast(m, syncType, options...)
}
