package commands

import (
	"fmt"
	"github.com/smowafy/asdf/internal/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"syscall"
)

var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "Sign in with account ID and password",
	//	Long: `echo long`,
	Args: cobra.ExactArgs(1),
	Run:  RunLoginCommand,
}

func RunLoginCommand(cmd *cobra.Command, args []string) {
	accountId := args[0]

	fmt.Printf("password:\n")

	password, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		panic(err)
	}

	asdfClient, err := client.NewAsdfClient(string(password), accountId)

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
