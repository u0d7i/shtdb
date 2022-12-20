package cmd

import (

    "shtdb/api/db"

	"github.com/spf13/cobra"
)

var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "db info",
	Long:  `A longer description of the subcommand goes here`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Info()
	},
}
