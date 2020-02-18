package cli

import (
	"bufio"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	gtypes "github.com/wlsaud619/hellochain/x/greeter/internal/types"
)

// greeting 모듈에 대한 부모 tx 명령을 반환
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greetingTxCmd := &cobra.Command {
		Use:				"greeter",
		Short:				"greeter transaction subcommands",
		DisableFlagParsing:		true,
		SuggestionsMinimumDistance:	2,
		RunE:				client.ValidateCmd,
	}

	greetingTxCmd.AddCommand(flags.PostCommands(
		GetCmdSayHello(cdc),
	) ... )

	return greetingTxCmd
}

// greeter 모듈에 대한 tx say 명령을 반환
func GetCmdSayHello(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		Use:	"say [body] [addr]",
		Short:	"send a greeting to another user. Usage: say [body] [address]",
		Args:	cobra.ExactArgs(2),
		RunE:	func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// --from 플래그에서 발신 계정 주소(sending account)를 가져온다
			sender := cliCtx.GetFromAddress()
			body := args[0]

			// 아규먼트에서 수신자 주소를 가져온다
			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())
			// greeting 메시지를 보내기 위한 Tx을 구성, sign, encode
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := gtypes.NewMsgGreet(sender, body, recipient)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// greeting 메시지가 포함된 tx를 build, sign 후 broadcast
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

}
