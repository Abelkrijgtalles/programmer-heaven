package main

import (
	"bufio"
	"fmt"
	"strings"
)

func input(p string, r *bufio.Reader) (string) {
	fmt.Print(p)
	var i, err = r.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(i)
}