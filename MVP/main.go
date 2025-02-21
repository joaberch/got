package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: got hello")
		return
	}

	if os.Args[1] == "hello" {
		fmt.Println("got got got")
		return
	}
	fmt.Println("Commande inconnue. Utilisez: got hello")
}
