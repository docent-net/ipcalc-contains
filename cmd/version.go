package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ipcalc-contains",
	Long:  `All software has versions. This is ipcalc-contains'es.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ipcalc-contains 0.1.0")
	},
}
