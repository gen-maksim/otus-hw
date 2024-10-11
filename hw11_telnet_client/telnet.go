package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	Conn    net.Conn
	closer  chan struct{}
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{in: in, out: out, timeout: timeout, address: address}
}

func (t *telnetClient) Connect() error {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	conn, err := dialer.DialContext(ctx, "tcp", t.address)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	t.closer = make(chan struct{})
	go func() {
		select {
		case <-t.closer:
			cancel()
			conn.Close()
		case <-ctx.Done():
			cancel()
			conn.Close()
		}
	}()
	t.Conn = conn
	return nil
}

func (t *telnetClient) Close() error {
	close(t.closer)
	return nil
}

func (t *telnetClient) Send() error {
	err := t.translate(t.in, t.Conn)
	if err != nil {
		return err
	}
	return nil
}

func (t *telnetClient) Receive() error {
	err := t.translate(t.Conn, t.out)
	if err != nil {
		return err
	}

	return nil
}

func (t *telnetClient) translate(r io.Reader, w io.Writer) (err error) {
	scanner := bufio.NewScanner(r)
	for {
		select {
		case <-t.closer:
			return nil
		default:
			if !scanner.Scan() {
				return nil
			}
			text := scanner.Text()
			w.Write([]byte(text + "\n"))
		}
	}
}
