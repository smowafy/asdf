package item

import(
	"github.com/spf13/cobra"
	"github.com/smowafy/asdf/internal/client"
)

var ItemSetCommand = &cobra.Command{
	Use:   "set",
	Short: "set item in a vault, sets in the default vault if no vault ID provided",
//	Long: `echo long`,
	Args: cobra.ExactArgs(2),
	Run: RunItemSetCommand,
}

// TODO: pass account ID and password only once to client
func RunItemSetCommand(cmd *cobra.Command, args []string) {
	asdfClient, err := client.NewAsdfClient(password, accountId)

	if err != nil {
		panic(err)
	}

	if err := asdfClient.SetItem(args[0], []byte(args[1])); err != nil {
		panic(err)
	}
}
