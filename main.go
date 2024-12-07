package main

import (
	"flag"
	"fmt"
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

}
