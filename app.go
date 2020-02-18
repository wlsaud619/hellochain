package hellochain

import (
	 sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/sdk-tutorials/hellochain/starter"

	// greeter의 types 가져오기 
	"github.com/wlsaud619/hellochain/x/greeter"
)

const appName = "hellochain"

var (
	// ModuleBasics는 app에 포함된 모든 모듈의 AppModuleBasic 구조체를 보유한다
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter		// helloChainApp extends starter.AppStarter

	greeterkey	*sdk.KVStoreKey	// greeter 모듈을 위한 store key
	greeterKeeper	greeter.Keeper	// greeter 모듈을 위한 keeper

}

// NewHelloChainApp은 완전히  구성된 SDK Application을 반환한다
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	// appName, tendermint의 logger, tendermint의 db를 사용하여 새로운 appStarter를 생성한다
	// ModuleBasicsManager에 포함시키기 위해 greeter의 AppModuleBasic을 전달
	appStarter := starter.NewAppStarter(appName, logger, db, greeter.AppModuleBasic{})

	// greeter의 store key를 만든다
	greeterKey := sdk.NewKVStoreKey(greeter.StoreKey)

	// keeper 생성
	greeterKeeper := greeter.NewKeeper(greeterKey, appStarter.Cdc)

	// starter 및 greeter로 app을 구성한다
	var app = &helloChainApp {
		appStarter,

		greeterKey,
		greeterKeeper,
	}

	// greeter의 완전한 AppModule을 ModuleManager에 추가한다
	greeterMod := greeter.NewAppModule(greeterKeeper)
	app.Mm.Modules[greeterMod.Name()] = greeterMod

	// main store에 greeter의 data를 위한 하위 공간을 만든다
	app.MountStore(greeterKey, sdk.StoreTypeDB)




	// 마지막 구성을 위해 초기화 한다 
	app.InitializeStarter()

	return app
}
