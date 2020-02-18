package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// 메시지와 querier를 greeter 모듈로 라우틩하는 데 사용
const RouterKey = "greeter"

// 메시지(MsgGreet) 정의
type MsgGreet struct {
	Body	string
	Sender	sdk.AccAddress
	Recipient	sdk.AccAddress
}

// 메시지(MsgGreeter)를 위한 생성자
func NewMsgGreet(sender sdk.AccAddress, body string, recipient sdk.AccAddress) MsgGreet {
	return MsgGreet {
		Body:		body,
		Sender:		sender,
		Recipient:	recipient,
	}
}

// 라우팅 할때 사용되는 Key 반환
func (msg MsgGreet) Route() string {
	return RouterKey
}

// 메시지의 Action 반환
func (msg MsgGreet) Type() string {
	return "greet"
}

// 메시지의 stateless를 검증
func (msg MsgGreet) ValidateBasic() error {
	// 메시지의 수신자가 비어 있다면
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Recipient.String())
	}
	if len(msg.Sender) == 0 || len(msg.Body) == 0 || len(msg.Recipient) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Sender, Recipient and/or Body cannot be empty")
	}
	return nil
}

// 메시지 서명에 필요한 주소를 반환
func (msg MsgGreet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// 메시지 서명을 위해 인코딩  
func (msg MsgGreet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

