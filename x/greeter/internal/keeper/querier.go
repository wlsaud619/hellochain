package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	greeter "github.com/wlsaud619/hellochain/x/greeter/internal/types"
)

// hellochain Querier가 지원하는 query endpoints
const(
	ListGreetings = "list"
)

// state 쿼리를 위한 module 수준의 라우터
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error)  {
		switch path[0] {
		case ListGreetings:
			return listGreetings(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unkown greeter query endpoint")
		}
	}
}

func listGreetings(ctx sdk.Context, params []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	greetingList := greeter.NewQueryResGreetings()

	// 현재 블록의 Context에 있는 값 중에 Greeting 구조의 Key:Value 값을 iterator에 저장함 
	iterator := keeper.GetGreetingsIterator(ctx)

	addr, err := sdk.AccAddressFromBech32(params[0])
	if err != nil {
		return nil, err
	}

	for ; iterator.Valid(); iterator.Next() {
		var greeting greeter.Greeting
		// iterator의 값(greeting)을 greeting 변수에 저장
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &greeting)

		// iterator의 값에서 얻은 주소와 검색하려는 주소(addr)가 같다면 주소를 Key로 갖는 맵(greetingList)에 추가
		if greeting.Recipient.Equals(addr) {
			greetingList[addr.String()] = append(greetingList[addr.String()], greeting)
		}
	}

	hellos, err := codec.MarshalJSONIndent(keeper.cdc, greetingList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return hellos, nil
}
