package greeter

import (
	"github.com/wlsaud619/hellochain/x/greeter/internal/keeper"
	"github.com/wlsaud619/hellochain/x/greeter/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouteKey = types.RouterKey
	StoreKey = types.StoreKey
)

var (
	NewKeeper = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier
	NewMsgGreet = types.NewMsgGreet
	NewGreeting = types.NewGreeting
)

type (
	Keeper		  = keeper.Keeper
	MsgGreet	  = types.MsgGreet
	Greeting	  = types.Greeting
	QueryResGreetings = types.QueryResGreetings
	GreetingsList	  = types.GreetingsList
)
