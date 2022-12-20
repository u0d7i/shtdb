package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "shtdb",
	Short: "shit db",
	Long:  `shelf/shit db`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute all
func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.app.yaml)")
	rootCmd.AddCommand(SubCmd)
	rootCmd.AddCommand(DbCmd)
}
