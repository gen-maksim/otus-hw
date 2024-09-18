package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	timeout int
)

func init() {
	flag.IntVar(&timeout, "timeout", 0, "timeout for connection")
}
func main() {
	flag.Parse()
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Too few arguments")
		return
	}

	url := args[1]
	port := args[2]
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
}
