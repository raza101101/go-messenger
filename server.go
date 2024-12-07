package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

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
	addClient(conn)
	defer removeClient(conn)
	//fmt.Println(conn.LocalAddr(), conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		broadcast(scanner.Text(), conn)
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

}

//

var (
	clients []net.Conn = make([]net.Conn, 0)
	mu                 = &sync.Mutex{}
)

func addClient(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clients = append(clients, conn)
}
func removeClient(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for index, client := range clients {
		if conn.RemoteAddr().String() == client.RemoteAddr().String() {
			clients = append(clients[:index], clients[index+1:]...)
		}
	}
}
func broadcast(message string, sender net.Conn) {

	mu.Lock()
	defer mu.Unlock()

	for _, client := range clients {
		if client.RemoteAddr().String() != sender.RemoteAddr().String() {
			fmt.Fprintf(client, "%s\n", message)
		}

	}

}
