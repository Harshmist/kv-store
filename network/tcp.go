package network

import (
	"bufio"
	"kv-store/nodes"
	"log"
	"net"
	//"strconv"
	"strings"
)

func StartTCP() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go Handler(conn)

	}

}

func Handler(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 1 {
			continue
		}
		var userRequest nodes.UserRequest
		value := strings.Join(fields[2:], " ")
		key := fields[1]
		node := nodes.Node(key)
		userRequest.Action = strings.ToUpper(fields[0])
		userRequest.Key = key
		userRequest.Value = value

		switch userRequest.Action {
		case "GET":

		}

		nodes.ChannelSlice[node] <- userRequest

	}

}
