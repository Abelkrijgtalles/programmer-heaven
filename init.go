package main

import (
	"flag"
	"fmt"
	"os"
)

type projectFile struct {
	projectName string
	projectDesc string
}

func getAndSaveDefaultProjectFile() {
	
	var pF = projectFile{
		projectName: "A awesome project",
		projectDesc: "The best project i have ever made.",
	}

	projectFile, err := os.Create("ph.json")
	if err != nil {
		panic(err)
	}
	projectFile.WriteString(fmt.Sprintf(`{"projectName":"%v","projectDesc":"%v"}`, pF.projectName, pF.projectDesc))
}

func handleInit(initCmd *flag.FlagSet, y *bool) {
	initCmd.Parse(os.Args[2:])
	
	if *y {
		getAndSaveDefaultProjectFile()
	} else {
		createProjectFile()
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
}