package main

import (
	"flag"
	"fmt"
	"os"
)

func handleAddApp(addCmd *flag.FlagSet, n *string) {
	addCmd.Parse(os.Args[2:])
	var name = *n

	_, err := os.Stat("apps")
	if os.IsNotExist(err) {
		fmt.Println("You need to run init first")
		os.Exit(1)
	}

	if name == "unnamed-app" {
		var questionName = input("Type y if you really dont have a name for your app. If you have a name type n: ")
		switch questionName {
		case "n":
			name = input("What is the app name: ")
		case "y":
		default:
			fmt.Println("That is not a valid option")
			os.Exit(1)
		}
	}

	var nameNumber = 0
	for {
		if nameNumber == 0 {
			_, err = os.Stat(fmt.Sprintf("apps/%v", name))
		} else {
			_, err = os.Stat(fmt.Sprintf("apps/%v-%v", name, nameNumber))
		}
		if os.IsNotExist(err) {
			break
		} else {
			nameNumber = nameNumber + 1
		}
	}
	if nameNumber > 0 {
		name = fmt.Sprintf("%v-%v", name, nameNumber)
	}

	err = os.Mkdir(fmt.Sprintf("apps/%v", name), 0750)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	fmt.Println("Added app:", name)
}
