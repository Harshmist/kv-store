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
	Action string
	Key    string
	Value  string
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
			switch action {
			case "DELETE":
				delete(Data, key)
			case "POST":
				Data[key] = value
				fmt.Println(Data)
				fmt.Println(ChannelSlice)
			case "SET":
				Data[key] = value
			}
		}
	}

}
