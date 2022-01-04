package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"

	"github.com/docent-net/ipcalc-contains/subnet"
)

var containsCmd = &cobra.Command{
	Use:   "contains",
	Short: "Check whether first subnet (1st argument) contains IP addr or other subnet (2nd argument)",
	Long:  `Check whether first subnet (1st argument) contains IP addr or other subnet (2nd argument)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getSubnetInfo(args))
	},
}

func init() {
	rootCmd.AddCommand(containsCmd)
}

func getSubnetInfo(args []string) string {
	if ! verifyArgs(args) {
		fmt.Println(`You must provide a subnet (slash notation, e.g.
192.168.0.0/24) and an IP address as exact 2 arguments for
this command, e.g: ipcalc-contains 192.168.0.0/24 192.168.100.1

Run --help for details`)
		os.Exit(1)
	}

	verificationString := ""

	if subnet.SubnetContains(args) {
		verificationString = fmt.Sprintf("CONTAINS! Network %s contains %s", args[0], args[1])
	} else {
		verificationString = fmt.Sprintf("NOPE! Network %s DOES NOT contain %s", args[0], args[1])
	}
	return verificationString
}

func verifyArgs(args []string) bool {
	if len(args) > 2 || len(args) == 0 {
		return false
	}

	return true
}