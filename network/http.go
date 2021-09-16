package network

import (
	"fmt"
	"kv-store/nodes"
	"net/http"
	"strconv"
	"strings"
)

func HttpHandleFuncs() {

	http.HandleFunc("/", UserAction)
	http.ListenAndServe(":8001", nil)

}

//HTTP HandleFuncs

func UserAction(w http.ResponseWriter, r *http.Request) {

	var userRequest nodes.UserRequest
	rtnChan := make(chan string)
	parts := strings.Split(r.URL.String(), "/")
	valueSlice := strings.Split(parts[3], "%20")
	value := strings.Join(valueSlice, " ")
	node := int(nodes.Node(parts[2]))

	userRequest.Action = parts[1]
	userRequest.Key = parts[2]
	if userRequest.Action == "list" {
		keyInt, err := strconv.Atoi(userRequest.Key)
		if err != nil {
			panic(err)
		}
		node = keyInt

	}
	userRequest.Value = value
	userRequest.RtnChan = rtnChan
	userRequest.Node = node

	nodes.ChannelSlice[node] <- userRequest
	httpMsgReceiver(rtnChan, w)
}

func httpMsgReceiver(ch chan string, w http.ResponseWriter) {
	{
		for {
			select {
			case printString := <-ch:
				fmt.Fprintf(w, printString)
				return
			}
		}
	}
}
