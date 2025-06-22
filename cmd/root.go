/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate secure, customizable passwords directly from your terminal",
	Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "Display help information",
		Long:  "Display detailed help information for the passgen command and its subcommands.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	})
}


