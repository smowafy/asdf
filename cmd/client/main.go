package main

import(
	/*
	"github.com/smowafy/asdf/internal/client"
	"os"
	"log"
	"github.com/spf13/cobra"
	*/
	"github.com/smowafy/asdf/cmd/client/commands"
)

func main() {
	/*
	c, err := client.NewAsdfClient(os.Args[1], "one")

	if err != nil {
		panic(err)
	}

	err = c.SignUp(os.Args[1], "one")

	if err != nil {
		panic(err)
	}
//	srpx, err := client.GetSrpKey(os.Args[1], "one")
	err = c.Ping()

	if err != nil {
		panic(err)
	}

	vault := make(client.Vault)

	vault["key"] = []byte("value")

	en, err := c.EncryptVault(vault)

	if err != nil { panic(err) }

	nv, err := c.DecryptVault(en)

	if err != nil { panic(err) }

	log.Printf("decrypted: %v, %v\n", nv, string(nv["key"]))
	*/
	commands.Execute()
}
