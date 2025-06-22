/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/fxckcode/passgen/utils"
	"github.com/spf13/cobra"
)

var (
	length    int
	symbols   bool
	numbers   bool
	uppercase bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate secure, customizable passwords directly from your terminal",
	Long: `Passgen is a powerful and flexible CLI tool for generating secure passwords.

You can customize the generated password by specifying the desired length and
choosing whether to include uppercase letters, numbers, and symbols.

Examples:
  generate                    # Generate a 12-character password (default)
  generate -l 16              # Generate a 16-character password
  generate -l 20 -u -n -s     # Generate a 20-character password with uppercase letters, numbers, and symbols

Ideal for developers, system admins, or anyone who needs strong and unique passwords on the fly.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, error := utils.GeneratePassword(length, symbols, numbers, uppercase)
		if error != nil {
			fmt.Println("Error: Password generation failed. Please check your criteria.")
		} else {
			fmt.Println("Encrypted Password:", password)
		}
	},
}

func init() {
	generateCmd.PersistentFlags().IntVarP(&length, "length", "l", 12, "Length of the password")
	generateCmd.PersistentFlags().BoolVarP(&symbols, "symbols", "s", false, "Include symbols in the password")
	generateCmd.PersistentFlags().BoolVarP(&numbers, "numbers", "n", false, "Include numbers in the password")
	generateCmd.PersistentFlags().BoolVarP(&uppercase, "uppercase", "u", false, "Include uppercase letters in the password")
	rootCmd.AddCommand(generateCmd)
}
