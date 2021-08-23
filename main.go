package main

import (
	"kv-store/network"
	"kv-store/store"
)

func main() {

	//Placeholder data for testing
	store.Data[1] = "Dan"
	store.Data[2] = "Sinead"

	go network.StartTCP()

	go store.RequestHandler()

	network.HttpHandleFuncs()

}
