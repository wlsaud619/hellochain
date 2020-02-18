package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	gtypes "github.com/wlsaud619/hellochain/x/greeter/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greeterQueryCmd := &cobra.Command {
		Use:				"greeter",
		Short:				"Querying commands for the greeter module",
		DisableFlagParsing:		true,
		SuggestionsMinimumDistance:	2,
		RunE:				client.ValidateCmd,
	}

	greeterQueryCmd.AddCommand(flags.GetCommands(
		GetCmdListGreetings(storeKey, cdc),
	) ...)

	return greeterQueryCmd
}

func GetCmdListGreetings(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		Use:	"list [addr]",
		Short:	"list greetings for address. Usage list [address]",
		Args:	cobra.ExactArgs(1),
		RunE:	func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			/*
				cliCtx.QueryWithData는 아래 설정된 특정 URL 형식을 사용하여 app을 쿼리한다
				route 문자열의 format: "/custom/greeter/list/addr"
				"/custom/"	추가한 custom module
				"/greeter/"	모듈의 QuerierRote
				"/list/"	인사말의 Querier(greeter's Querier) 의 특정 endpoint
				"/addr"		쿼리의 파라미터

			*/

			route := fmt.Sprintf("custom/%s/list/%s", queryRoute, addr)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return nil
			}

			out := gtypes.NewQueryResGreetings()
			cdc.MustUnmarshalJSON(res, &out)


			fmt.Println("cliCtx: ", cliCtx)
			fmt.Println("args: ", args)

			return cliCtx.PrintOutput(out)
		},
	}
}
