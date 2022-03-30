package main

import (
	"flag"
	"fmt"
	"os"
)

func handleAddApp(addCmd *flag.FlagSet, n *string) {
	addCmd.Parse(os.Args[2:])
	var name = *n

	if name == "unnamed-app" {
		name = input("Type y if you really dont have a name for your app. If you have a name type n: ")
		switch name {
		case "n":
			name = input("What is the app name: ")
		case "y":
			name = "unnamed-app"
		default:
			fmt.Println("That is not a valid option")
			os.Exit(1)
		}
	}

	_, err := os.Stat("apps")
	if os.IsNotExist(err) {
		fmt.Println("You need to run init first")
		os.Exit(1)
	}

	err = os.Mkdir(fmt.Sprintf("apps/%v", name), 0750)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	if os.IsExist(err) {
		fmt.Println("App with that name already exists")
		os.Exit(1)
	}
	fmt.Println("Added app:", name)
}
