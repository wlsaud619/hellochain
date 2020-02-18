package greeter

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/wlsaud619/hellochain/x/greeter/client/cli"
	gtypes "github.com/wlsaud619/hellochain/x/greeter/internal/types"
)

// AppModuleBasic은 Module의 최소 구조체
type AppModuleBasic struct {
	starter.BlankModuleBasic
}

// AppModule에 전체 Module이 포함된다
type AppModule struct {
	starter.BlankModule
	keeper     Keeper
	ModuleName string
}

// 인터페이스가 올바로 구분되어 있는지 확인
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// RegisterCodec은 encoding/decoding을 위한 모듈 메시지를 등록한다
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(gtypes.MsgGreet{}, "greeter/SayHello", nil)
}

// NewHandler는 메시지를 적절한 handler 함수로 라우팅
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// 요청된 쿼리를 올바른 querier로 라우팅
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// 쿼리를 모듈에 라우팅하기 위해 사용됨
func (am AppModule) QuerierRoute() string {
	return am.ModuleName
}

// 모듈에서 지원하는 모든 CLI 명령어와 쿼리를 모아서 반환
func (ab AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(gtypes.StoreKey, cdc)
}

// 모듈에서 지원하는 모든 CLI 명령어와 쿼리를 모아서 반환
func (ab AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(gtypes.StoreKey, cdc)
}

// 이 모듈의 AppModule 전체 구조체를 구성
func NewAppModule(keeper Keeper) AppModule {
	blank := starter.NewBlankModule(gtypes.ModuleName, keeper)
	return AppModule{blank, keeper, gtypes.ModuleName}
}
