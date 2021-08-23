package store

import (
	"log"
	"strconv"
)

var (
	Data        = make(map[int]string)
	PostChannel = make(chan string)
	SetChannel  = make(chan [2]string)
	DelChannel  = make(chan int)
)

func RequestHandler() {
	for {
		select {
		case val1 := <-PostChannel:
			Data[len(Data)+1] = val1

		case val2 := <-SetChannel:
			key, err := strconv.Atoi(val2[0])
			if err != nil {
				log.Fatal("Fatal error")
			}
			Data[key] = val2[1]

		case delKey := <-DelChannel:
			Data[delKey] = ""

		}
	}
}
