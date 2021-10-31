package item

import (
	"github.com/spf13/cobra"
)

var (
	sessionKey     string
	passwordLength int
)

var ItemCommand = &cobra.Command{
	Use:   "item",
	Short: "item commands",
}

func init() {
	ItemCommand.AddCommand(ItemGetCommand)
	ItemCommand.AddCommand(ItemSetCommand)

	ItemCommand.PersistentFlags().StringVar(&sessionKey, "session-key", "", "Session key")
}
