package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gtypes "github.com/wlsaud619/hellochain/x/greeter/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// 데이터 저장소에 대한 연결을 유지하고 여러 method에 대한 getter/setter를 제공한다
// State machine의 일부이다 
type Keeper struct {
	storeKey	sdk.StoreKey	// sdk.Context에 접근하기 위한 노투출되지 않은 Key
	cdc		*codec.Codec	// binary encoding/decoding을 위한 wire codec
}

// greeter Keeper에 대한 새로운 인스턴스 생성
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper {
		storeKey:	storeKey,
		cdc:		cdc,
	}
}

// 주어진 주소와 보낸 사람의 인사말을 반환 
func (k Keeper) GetGreetings(ctx sdk.Context, addr sdk.AccAddress, from sdk.Address) gtypes.GreetingsList {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr)) {
		return gtypes.GreetingsList{}
	}
	bz := store.Get([]byte(addr))
	var list gtypes.GreetingsList
	k.cdc.MustUnmarshalBinaryBare(bz, &list)

	if from != nil {
		//
		var fromList gtypes.GreetingsList
		for _, g := range list {
			if g.Sender.Equals(from) {
				fromList = append(fromList, g)
			}
		}
		return fromList
	}
	return list
}

// 주어진 주소에 대한 인사말을 저장
func (k Keeper) SetGreeting(ctx sdk.Context, greeting gtypes.Greeting) {
	if greeting.Sender.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	list := k.GetGreetings(ctx, greeting.Recipient, nil)
	list = append(list, greeting)
	store.Set(greeting.Recipient.Bytes(), k.cdc.MustMarshalBinaryBare(list))
}

// "Key:Value"가 "주소:인사말들"인 모든 이름에 대해 반복자(iterator)를 반환 
func (k Keeper) GetGreetingsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
