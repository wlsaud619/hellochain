package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/wlsaud619/hellochain"
	"github.com/wlsaud619/hellochain/starter"
)

func main() {

	starter.BuildModuleBasics()

	rootCmd := starter.NewCLICommand()

	txCmd := starter.TxCmd(starter.Cdc)
	queryCmd := starter.QueryCmd(starter.Cdc)

	// Tx, Query 명령어 추가 
	app.ModuleBasics.AddTxCommands(txCmd, starter.Cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, starter.Cdc)
	rootCmd.AddCommand(txCmd, QueryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HC", starter.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
