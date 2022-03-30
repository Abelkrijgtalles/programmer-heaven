package main

import (
	"flag"
	"fmt"
	"os"
)

func getAndSaveDefaultProjectFile() {
	var projectName = "A awesome project"
	var projectDesc = "The best project i have ever made."

	projectFile, err := os.Create("ph.json")
	if err != nil {
		panic(err)
	}
	projectFile.WriteString(fmt.Sprintf(`{"projectName":"%v","projectDesc":"%v"}`, projectName, projectDesc))

	fmt.Println("Default project file created")
}

func handleInit(initCmd *flag.FlagSet, y *bool) {
	initCmd.Parse(os.Args[2:])
	
	if *y {
		getAndSaveDefaultProjectFile()
	} else {
		createProjectFile()
	}

	err := os.Mkdir("apps", 0750)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir("shared", 0750)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func createProjectFile() {
	var projectName = input("What is the project name: ")
	var projectDesc = input("What is the project description: ")

	projectFile, err := os.Create("ph.json")
	if err != nil {
		panic(err)
	}
	projectFile.WriteString(fmt.Sprintf(`{"projectName":"%v","projectDesc":"%v"}`, projectName, projectDesc))

	fmt.Println("Project file created")
}