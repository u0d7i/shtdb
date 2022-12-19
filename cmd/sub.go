package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SubCmd = &cobra.Command{
	Use:   "sub",
	Short: "Run the subcommand",
	Long:  `A longer description of the subcommand goes here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the subcommand!")
	},
}
