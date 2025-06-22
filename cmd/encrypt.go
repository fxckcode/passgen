/*
Copyright © 2025 fxckcode

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/fxckcode/passgen/utils"
)

var (
	pass string
	cost int = 11
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a password using bcrypt",
	Long: `Encrypt a plaintext password using the bcrypt hashing algorithm.

You can specify the password via the --password flag and optionally set the bcrypt cost (default is 11).

Examples:
  passgen encrypt -p mypassword
  passgen encrypt --password "supersecret" --cost 12

The cost defines the computational complexity of the bcrypt hashing process. Higher values are more secure but slower.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.EncryptPassword(pass, cost)
		if err != nil {
			fmt.Println("Error: Password encryption failed. Please check your input.")
			return
		}
		fmt.Println("Encrypted Password:", password)
	},
}

func init() {
	encryptCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "Password to encrypt")
	encryptCmd.PersistentFlags().IntVarP(&cost, "cost", "c", 11, "Bcrypt hash cost (default is 11)")
	encryptCmd.MarkPersistentFlagRequired("password")
	rootCmd.AddCommand(encryptCmd)
}
