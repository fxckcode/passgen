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
	password string
	hash string
)

// checkpassCmd represents the checkpass command
var checkpassCmd = &cobra.Command{
	Use:   "checkpass",
	Short: "Verify if a password matches a given bcrypt hash",
	Long: `Check if a plaintext password matches a bcrypt hash.

This is useful for verifying stored password hashes, especially in login or authentication systems.

Examples:
  passgen checkpass -p mypassword -H '$2a$11$abc...xyz'
  passgen checkpass --password "admin123" --hash "$2b$12$abcdef..."`,
	Run: func(cmd *cobra.Command, args []string) {
		ok, err := utils.CheckPassword(password, hash)
		if err != nil {
			fmt.Printf("Error checking password: %v\n", err)
			return
		}
		if ok {
			fmt.Println("Password is valid.")
		} else {
			fmt.Println("Password is invalid.")
		}
	},
}

func init() {
	checkpassCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password to check")
	checkpassCmd.PersistentFlags().StringVarP(&hash, "hash", "H", "", "Bcrypt hash to compare against")
	checkpassCmd.MarkPersistentFlagRequired("password")
	checkpassCmd.MarkPersistentFlagRequired("hash")
	rootCmd.AddCommand(checkpassCmd)
}
