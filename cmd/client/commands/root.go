package commands

import(
	"github.com/spf13/cobra"
	"github.com/smowafy/asdf/cmd/client/commands/item"
)

var (
	rootCmd = &cobra.Command{Use: "asdf"}
)

func init() {
	rootCmd.AddCommand(SignUpCommand)
	rootCmd.AddCommand(item.ItemCommand)
}

func Execute() error {
	return rootCmd.Execute()
}
