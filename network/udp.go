package network

import (
	"fmt"
	"kv-store/nodes"
	"net"
	"strconv"
	"strings"
)

func StartUDP() {

	service := "localhost:8002"

	s, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			return
		}

		input := strings.TrimSpace(string(buffer[:n]))
		fields := strings.Split(string(input), " ")
		var userRequest nodes.UserRequest
		value := strings.Join(fields[2:], " ")
		key := fields[1]
		node := int(nodes.Node(key))

		rtnChan := make(chan string)

		userRequest.Action = strings.ToUpper(fields[0])
		userRequest.Key = key
		if userRequest.Action == "LIST" {
			keyInt, err := strconv.Atoi(key)
			if err != nil {
				panic(err)
			}
			node = keyInt

		}
		userRequest.Value = value
		userRequest.RtnChan = rtnChan
		userRequest.Node = node

		nodes.ChannelSlice[node] <- userRequest
		UDPMsgReceiver(rtnChan, addr, connection)
	}
}

func UDPMsgReceiver(ch chan string, addr *net.UDPAddr, conn *net.UDPConn) {
	for {
		select {
		case printString := <-ch:
			conn.WriteToUDP([]byte(printString), addr)
			return
		}
	}

}
