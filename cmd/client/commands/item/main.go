package item

import (
	"github.com/spf13/cobra"
)

var (
	accountId string
	password  string
)

var ItemCommand = &cobra.Command{
	Use:   "item",
	Short: "item commands",
}

func init() {
	ItemCommand.AddCommand(ItemGetCommand)
	ItemCommand.AddCommand(ItemSetCommand)

	ItemCommand.PersistentFlags().StringVar(&accountId, "id", "", "account ID")
	ItemCommand.PersistentFlags().StringVar(&password, "password", "", "password")

}
