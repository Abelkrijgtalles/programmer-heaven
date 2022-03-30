package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	var initCmdYes = initCmd.Bool("y", false, "Save as default")
	addCmdApp := flag.NewFlagSet("add-app", flag.ExitOnError)
	addCmdAppName := addCmdApp.String("name", "unnamed-app", "Add an app")

	if len(os.Args) < 2 {
		fmt.Println("Expected to get a subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		handleInit(initCmd, initCmdYes)
	case "add-app":
		handleAddApp(addCmdApp, addCmdAppName)
	default:
		fmt.Println("That is not a valid subcommand")
	}
}
