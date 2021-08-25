package main

import (
	"fmt"
	"kv-store/network"
	"kv-store/store"
)

func main() {
	go network.StartTCP()
	go store.RequestHandler()
	go network.HttpHandleFuncs()
	go network.StartUDP()

	fmt.Scanln()
}
