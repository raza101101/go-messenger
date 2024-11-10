package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	mode string
	nick string
)

func main() {

	flag.StringVar(&mode, "mode", "", "what mode would you like")
	flag.StringVar(&nick, "nick", "", "enter your nickname")

	flag.Parse()

	if nick == "" && mode == "client" {
		fmt.Println("error please enter a nickname")
		return
	}

	if mode == "" {
		fmt.Println("enter a mode.")
		return
	}
	if mode == "server" {
		fmt.Println("server mode")
		err := server()
		if err != nil {
			fmt.Println(err)
			return
		}

	} else if mode == "client" {
		fmt.Println("client mode")
		err := client()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("error, provide a valid mode (server, client)")
		return
	}

	//fmt.Print("Enter text: ") practice shi
	//var input string
	//fmt.Scanln(&input)
	//fmt.Print(input)
}
func server() error {

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

}
func client() error {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}
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
