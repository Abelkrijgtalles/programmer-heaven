package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	var initCmdYes = initCmd.Bool("y", false, "Save as default")

	if len(os.Args) < 2 {
		fmt.Println("Expected to get a subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		handleInit(initCmd, initCmdYes)
	default:
		fmt.Println("That is not a valid subcommand")
	}
}
