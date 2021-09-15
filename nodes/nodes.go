package nodes

import (
	"fmt"
	"hash/fnv"
	"strings"
)

var (
	NumOfNodes   uint64
	ChannelSlice = make([]chan UserRequest, 0, 10)
)

type UserRequest struct {
	Action  string
	Key     string
	Value   string
	RtnChan chan string
	Node    int
}

//Hash Function

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func Node(s string) uint64 {
	return Hash(s) % NumOfNodes
}

func Nodes(ch chan UserRequest) {
	Data := make(map[string]string)

	for {
		select {
		case request := <-ch:
			action := strings.ToUpper(request.Action)
			key := request.Key
			value := request.Value
			node := request.Node
			switch action {
			case "DELETE":
				delete(Data, key)
				rtnMsg := fmt.Sprintf("%v deleted from node %v\n", key, node)
				request.RtnChan <- rtnMsg
			case "POST":
				Data[key] = value
				rtnMsg := fmt.Sprintf("Key: %v created and set to: %v\n", key, value)
				request.RtnChan <- rtnMsg
			case "SET":
				Data[key] = value
				rtnMsg := fmt.Sprintf("Key :%v now set to: %v\n", key, value)
				request.RtnChan <- rtnMsg

			case "GET":
				key := request.Key
				request.RtnChan <- fmt.Sprintf("%v: %v\n", key, Data[key])
			case "LIST":
				printString := "List of Data:\n"
				for k, v := range Data {
					newString := fmt.Sprintf("%v: %v\n", k, v)
					printString = printString + newString
				}
				request.RtnChan <- printString
			}
		}
	}

}
