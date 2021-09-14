package main

import (
	"fmt"
	"kv-store/network"
	"kv-store/nodes"
)

func main() {
	go nodes.Node0()
	go nodes.Node1()
	go nodes.Node2()
	go nodes.Node3()
	go network.HttpHandleFuncs()
	go network.StartTCP()
	//go network.StartUDP()

	fmt.Scanln()
}
