package cmd

import (
	"fmt"
	"os"
	"net"
	"github.com/spf13/cobra"
)

var containsCmd = &cobra.Command{
	Use:   "contains",
	Short: "Check whether first subnet (1st argument) contains IP addr (2nd argument)",
	Long:  `Check whether first subnet (1st argument) contains IP addr (2nd argument)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(subnetContains(args))
	},
}

func init() {
	rootCmd.AddCommand(containsCmd)
}

func subnetContains(args []string) string {
	if ! verifyArgs(args) {
		fmt.Println(`You must provide a subnet (slash notation, e.g.
192.168.0.0/24) and an IP address as exact 2 arguments for
this command, e.g: ipcalc-contains 192.168.0.0/24 192.168.100.1

Run --help for details`)
		os.Exit(1)
	}

	_, ipv4Net, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ipv4Addr := net.ParseIP(args[1])
	if ipv4Addr == nil {
		fmt.Println(string(args[1]) + " is not recognized as IPv4 address")
		os.Exit(1)
	}

	verificationString := ""

	if ipv4Net.Contains(ipv4Addr) {
		verificationString = fmt.Sprintf("CONTAINS! Network %s contains IP %s", ipv4Net, ipv4Addr)
	} else {
		verificationString = fmt.Sprintf("NOPE! Network %s DOES NOT contain %s", ipv4Net, ipv4Addr)
	}
	return verificationString
}

func verifyArgs(args []string) bool {
	if len(args) > 2 || len(args) == 0 {
		return false
	}

	return true
}