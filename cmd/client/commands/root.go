package commands

import (
	"github.com/smowafy/asdf/cmd/client/commands/item"
	"github.com/spf13/cobra"
)

var (
	passwordLength int
	rootCmd        = &cobra.Command{Use: "asdf"}
)

func init() {
	rootCmd.AddCommand(SignUpCommand)
	rootCmd.AddCommand(LoginCommand)
	rootCmd.AddCommand(RandCommand)
	rootCmd.AddCommand(item.ItemCommand)

	RandCommand.Flags().IntVarP(&passwordLength, "length", "l", 15, "Password length")
}

func Execute() error {
	return rootCmd.Execute()
}
