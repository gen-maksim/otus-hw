package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var timeout string

func init() {
	flag.StringVar(&timeout, "timeout", "30s", "timeout for connection")
}

func main() {
	flag.Parse()
	args := os.Args
	if len(args) < 3 {
		log.Println("Too few arguments")
		return
	}

	url := args[2]
	port := args[3]
	in := os.Stdin
	out := os.Stdout
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Println("Invalid timeout")
		return
	}

	c := NewTelnetClient(fmt.Sprintf("%v:%v", url, port), duration, in, out)
	c.Connect()
	defer c.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.Receive()
	}()

	go func() {
		defer wg.Done()
		c.Send()
	}()

	wg.Wait()
}
