package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(p string) (string) {
	var r = bufio.NewReader(os.Stdin)
	fmt.Print(p)
	var i, err = r.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(i)
}