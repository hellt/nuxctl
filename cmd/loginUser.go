package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CredFPath is a path to the user credentials file
var CredFPath string

func loginUser(cmd *cobra.Command, args []string) {
	user.LoadCredentials(CredFPath)
	_, err := user.Login()
	if err != nil {
		fmt.Println(err)
	}
}
