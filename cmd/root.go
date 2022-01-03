package cmd

import (
	"fmt"
	"os"

	// "github.com/docent-net/ipcalc-contains"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ipcalc-contains",
		Short: "check whether subnet contains IP address",
		Long:  `check whether subnet contains IP address`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

// func init() {
// 	rootCmd.AddCommand(versionCmd)
//   }

// execute cmd with default values as defined in
// https://github.com/spf13/cobra/issues/823#issuecomment-949732548
func Execute(defCmd string) {
	var cmdFound bool
  	cmd :=rootCmd.Commands()

	  for _,a:=range cmd{
		for _,b:=range os.Args[1:] {
		  if a.Name()==b {
		   cmdFound=true
			break
		  }
		}
	  }
	  if cmdFound == false {
		args:=append([]string{defCmd}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	  }
	  if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }
}
