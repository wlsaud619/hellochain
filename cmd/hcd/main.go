package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/wlsaud619/hellochain"
	"github.com/wlsaud619/hellochain/starter"
)

func main() {

	params := starter.NewServerCommandParams (
		"hcd"  // 명렁어 이름
		"hellochain AppDaemon"  // 설명
		starter.NewApp.Creator(app.NewHelloChainApp),  // 앱을 구성하기 위한 method
		starter.NewAppExporter(app.NewHelloChainApp),  // chain state를 내보내기 위한 method
	)

	serverCmd := starter.NewServerCommand(params)

	// flags 추가 및 준비 
	executor := cli.PrepareBaseCmd(serverCmd, "HC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}

}
