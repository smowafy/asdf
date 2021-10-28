package item

import (
	"fmt"
	"github.com/smowafy/asdf/internal/client"
	"github.com/spf13/cobra"
)

var ItemGetCommand = &cobra.Command{
	Use:   "get",
	Short: "get item from a vault, gets from the default vault if no vault ID provided",
	//	Long: `echo long`,
	Args: cobra.ExactArgs(1),
	Run:  RunItemGetCommand,
}

// TODO: pass account ID and password only once to client
func RunItemGetCommand(cmd *cobra.Command, args []string) {
	asdfClient, err := client.LoadSession(sessionKey)

	if err != nil {
		panic(err)
	}

	item, err := asdfClient.GetItem(args[0])

	if err != nil {
		panic(err)
	}

	fmt.Printf("value: %v\n", string(item))
}
