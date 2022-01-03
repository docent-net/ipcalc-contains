package main

import (
	"github.com/docent-net/ipcalc-contains/cmd"
)

func main() {
	// set a default command
	// https://github.com/spf13/cobra/issues/823#issuecomment-949732548
	defCmd := "contains"
  	cmd.Execute(defCmd)
}
