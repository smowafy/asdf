package commands

import (
	"github.com/smowafy/asdf/internal/client"
	"github.com/spf13/cobra"
)

var SignUpCommand = &cobra.Command{
	Use:   "signup",
	Short: "sign up with a new account ID",
	//	Long: `echo long`,
	Args: cobra.ExactArgs(2),
	Run:  RunSignUpCommand,
}

// TODO: pass account ID and password only once to client
func RunSignUpCommand(cmd *cobra.Command, args []string) {
	accountId, password := args[0], args[1]
	asdfClient, err := client.NewAsdfClient(password, accountId)

	if err != nil {
		panic(err)
	}

	if err = asdfClient.SignUp(password, accountId); err != nil {
		panic(err)
	}
}
