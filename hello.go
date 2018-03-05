package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	args := strings.Join(os.Args[1:]," ")
	fmt.Println(ToUpper(args))
}

func ToUpper(s string) (u string) {
	return strings.ToUpper(s)
}