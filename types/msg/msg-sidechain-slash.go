package msg

import (
	"bytes"
	"fmt"

	asc "github.com/aximchain/asc-double-sign-sdk/types/asc"
	"github.com/aximchain/go-sdk/common/types"
	gosdkmsg "github.com/aximchain/go-sdk/types/msg"
)

const (
	TypeSideChainSubmitEvidence = "asc_submit_evidence"

	SideChainSlashMsgRoute = "slashing"
)

type MsgAscSubmitEvidence struct {
	Submitter types.AccAddress `json:"submitter"`
	Headers   []*asc.Header    `json:"headers"`
}

func NewMsgAscSubmitEvidence(submitter types.AccAddress, headers []*asc.Header) MsgAscSubmitEvidence {

	return MsgAscSubmitEvidence{
		Submitter: submitter,
		Headers:   headers,
	}
}

func (MsgAscSubmitEvidence) Route() string {
	return SideChainSlashMsgRoute
}

func (MsgAscSubmitEvidence) Type() string {
	return TypeSideChainSubmitEvidence
}

func headerEmptyCheck(header *asc.Header) error {
	if header.Number == 0 {
		return fmt.Errorf("header number can not be zero ")
	}
	if header.Difficulty == 0 {
		return fmt.Errorf("header difficulty can not be zero")
	}
	if header.Extra == nil {
		return fmt.Errorf("header extra can not be empty")
	}

	return nil
}

func (msg MsgAscSubmitEvidence) ValidateBasic() error {
	if len(msg.Submitter) != types.AddrLen {
		return fmt.Errorf("Expected delegator address length is %d, actual length is %d", types.AddrLen, len(msg.Submitter))
	}

	if err := headerEmptyCheck(msg.Headers[0]); err != nil {
		return err
	}
	if err := headerEmptyCheck(msg.Headers[1]); err != nil {
		return err
	}
	if msg.Headers[0].Number != msg.Headers[1].Number {
		return fmt.Errorf("The numbers of two block headers are not the same")
	}
	if !bytes.Equal(msg.Headers[0].ParentHash[:], msg.Headers[1].ParentHash[:]) {
		return fmt.Errorf("The parent hash of two block headers are not the same")
	}
	signature1, err := msg.Headers[0].GetSignature()
	if err != nil {
		return fmt.Errorf("Failed to get signature from block header, %s", err.Error())
	}
	signature2, err := msg.Headers[1].GetSignature()
	if err != nil {
		return fmt.Errorf("Failed to get signature from block header, %s", err.Error())
	}
	if bytes.Compare(signature1, signature2) == 0 {
		return fmt.Errorf("The two blocks are the same")
	}

	return nil
}

func (msg MsgAscSubmitEvidence) GetSignBytes() []byte {
	bz := gosdkmsg.MsgCdc.MustMarshalJSON(msg)
	return gosdkmsg.MustSortJSON(bz)
}

func (msg MsgAscSubmitEvidence) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Submitter}
}

func (msg MsgAscSubmitEvidence) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
