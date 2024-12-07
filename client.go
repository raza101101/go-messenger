package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() error {
	conn, err := net.Dial("tcp", "100.71.120.116:8080")
	if err != nil {
		return err
	}
	go incoming(conn)
	//fmt.Fprintf(conn, "this is a message being sent from client to the server via tcp if you are reading this on the server side this has worked\n")
	//return nil
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(conn, "[%s] %s\n", nick, scanner.Text())

	}

	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil

}
func incoming(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())

	}

	if scanner.Err() != nil {
		fmt.Println("error", scanner.Err())
	}

}
