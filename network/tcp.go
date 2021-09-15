package network

import (
	"bufio"
	"io"
	"kv-store/nodes"
	"log"
	"net"
	"strconv"

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

		go msgReciever(rtnChan, conn)

		nodes.ChannelSlice[node] <- userRequest
	}

}

func msgReciever(ch chan string, conn net.Conn) {
	go func() {
		for {
			select {
			case printString := <-ch:
				io.WriteString(conn, printString)
				return
			}
		}
	}()
}
