/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
