package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var timeout string

func init() {
	flag.StringVar(&timeout, "timeout", "10s", "timeout for connection")
}

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Println("Invalid timeout")
		return
	}

	args := flag.Args()
	if len(args) < 2 {
		log.Println("Too few arguments")
		return
	}

	url := args[0]
	port := args[1]
	in := os.Stdin
	out := os.Stdout

	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt, os.Signal(syscall.SIGTERM))

	c := NewTelnetClient(fmt.Sprintf("%v:%v", url, port), duration, in, out)
	c.Connect()
	go func() {
		<-kill
		c.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
			kill <- os.Interrupt
		}()
		c.Receive()
	}()

	go func() {
		defer func() {
			wg.Done()
			kill <- os.Interrupt
		}()
		c.Send()
	}()

	wg.Wait()
}
