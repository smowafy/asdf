package commands

import (
	"github.com/smowafy/asdf/internal/client"
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"golang.org/x/term"
	"syscall"
)

var SignUpCommand = &cobra.Command{
	Use:   "signup [ACCOUNT_ID]",
	Short: "sign up with a new account ID",
	//	Long: `echo long`,
	Args: cobra.ExactArgs(2),
	Run:  RunSignUpCommand,
}

func RunSignUpCommand(cmd *cobra.Command, args []string) {
	accountId := args[0]

	fmt.Printf("password:\n")

	password, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		panic(err)
	}

	fmt.Printf("password confirmation:\n")

	passwordConfirmation, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		panic(err)
	}

	passwordString := string(password)
	passwordConfirmationString := string(passwordConfirmation)

	if passwordString != passwordConfirmationString {
		fmt.Printf("Passwords do not match\n")
		os.Exit(1)
	}

	asdfClient, err := client.NewAsdfClient(passwordString, accountId)

	if err != nil {
		panic(err)
	}

	if err = asdfClient.SignUp(passwordString, accountId); err != nil {
		panic(err)
	}
}
