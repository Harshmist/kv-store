package main

import (
	"flag"
	"fmt"
	"kv-store/network"
	"kv-store/nodes"
)

var (
	numOfNodes = flag.Uint64("nodes", 4, "Number of nodes to be created")
)

func main() {
	flag.Parse()
	nodes.NumOfNodes = *numOfNodes
	for i := 0; i < int(*numOfNodes); i++ {
		channel := make(chan nodes.UserRequest)
		nodes.ChannelSlice = append(nodes.ChannelSlice, channel)
		go nodes.Nodes(channel)

	}

	go network.HttpHandleFuncs()
	go network.StartTCP()
	//go network.StartUDP()

	fmt.Println(*numOfNodes)

	fmt.Scanln()
}
