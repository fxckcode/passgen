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
	pass string
	cost int = 11
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
