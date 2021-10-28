package commands

import (
	"fmt"
	"github.com/smowafy/asdf/internal/client"
	"github.com/spf13/cobra"
)

var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "Sign in with account ID and password",
	//	Long: `echo long`,
	Args: cobra.ExactArgs(2),
	Run:  RunLoginCommand,
}

// TODO: pass account ID and password only once to client
func RunLoginCommand(cmd *cobra.Command, args []string) {
	accountId, password := args[0], args[1]
	asdfClient, err := client.NewAsdfClient(password, accountId)

	if err != nil {
		panic(err)
	}

	// TODO check against the DB, for now it's just sessions
	sessionKey, err := client.SaveSession(*asdfClient)

	if err != nil {
		panic(err)
	}

	fmt.Printf("session key:\n%s\n", sessionKey)
}
