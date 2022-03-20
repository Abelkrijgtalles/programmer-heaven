package main

import (
	"flag"
	"io"
	"net/http"
	"os"
)

func getAndSaveDefaultProjectFile() {
	resp, err := http.Get("https://raw.githubusercontent.com/Abelkrijgtalles/programmer-heaven/main/downloadable/ph.json")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	projectFile, err := os.Create("ph.json")
	if err != nil {
		panic(err)
	}
	projectFile.WriteString(string(body))
}

func handleInit(initCmd *flag.FlagSet, y *bool) {
	initCmd.Parse(os.Args[2:])
	
	if *y {
		getAndSaveDefaultProjectFile()
	}
}
