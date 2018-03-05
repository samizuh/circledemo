package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	args := strings.Join(os.Args[1:]," ")
	fmt.Println(toUpper(args))
}

func toUpper(s string) (u string) {
	return strings.ToUpper(s)
}