package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/wlsaud619/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/wlsaud619/hellochain/x/greeter"
)

func main() {

	// starter.BuildModuleBasics()
	starter.BuildModuleBasics(greeter.AppModuleBasic{})

	rootCmd := starter.NewCLICommand()

	txCmd := starter.TxCmd(starter.Cdc)
	queryCmd := starter.QueryCmd(starter.Cdc)

	// Tx, Query 명령어 추가 
	app.ModuleBasics.AddTxCommands(txCmd, starter.Cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, starter.Cdc)
	rootCmd.AddCommand(txCmd, queryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HC", starter.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
