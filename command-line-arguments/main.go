package main

import (
	"fmt"
	"os"
)

// go build command-line-arguments/main.go
// ./main a b c d

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	args := os.Args[3]

	fmt.Println(argsWithProg)    // [./main a b c d]
	fmt.Println(argsWithoutProg) // [a b c d]
	fmt.Println(args)            // c
}
