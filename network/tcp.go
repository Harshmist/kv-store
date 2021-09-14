package network

import (
	"bufio"
	"fmt"
	"io"
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

		switch fields[0] {

		case "POST":

			if len(fields) < 3 {
				io.WriteString(conn, "No value added! \n")
			}
			var sendSlice = make([]string, 2, 2)
			key := fields[1]
			value := strings.Join(fields[2:], " ")
			node := nodes.Node(key)
			sendSlice[0] = key
			sendSlice[1] = value
			switch node {
			case 0:
				nodes.Node0Post <- sendSlice
				io.WriteString(conn, value+" added!\n")
			case 1:
				nodes.Node1Post <- sendSlice
				io.WriteString(conn, value+" added!\n")
			case 2:
				nodes.Node2Post <- sendSlice
				io.WriteString(conn, value+" added!\n")
			case 3:
				nodes.Node3Post <- sendSlice
				io.WriteString(conn, value+" added!\n")
			}

		case "LIST":
			node := fields[1]

			switch node {
			case "0":
				io.WriteString(conn, fmt.Sprintf("All data from node %v\n", node))
				for k, v := range nodes.Data0 {
					io.WriteString(conn, fmt.Sprintf("%v: %v\n", k, v))
				}
			case "1":
				io.WriteString(conn, fmt.Sprintf("All data from node %v\n", node))
				for k, v := range nodes.Data1 {
					io.WriteString(conn, fmt.Sprintf("%v: %v\n", k, v))
				}
			case "2":
				io.WriteString(conn, fmt.Sprintf("All data from node %v\n", node))
				for k, v := range nodes.Data2 {
					io.WriteString(conn, fmt.Sprintf("%v: %v\n", k, v))
				}
			case "3":
				io.WriteString(conn, fmt.Sprintf("All data from node %v\n", node))
				for k, v := range nodes.Data3 {
					io.WriteString(conn, fmt.Sprintf("%v: %v\n", k, v))
				}
			default:
				io.WriteString(conn, "Node does not exist!")

			}

		case "SET":
			if len(fields) < 3 {
				io.WriteString(conn, "Format should be <string Key> <string Value> \n")
			}
			var sendSlice = make([]string, 2, 2)
			key := fields[1]
			value := strings.Join(fields[2:], " ")
			node := nodes.Node(key)
			sendSlice[0] = key
			sendSlice[1] = value

			switch node {
			case 0:
				nodes.Node0Post <- sendSlice
			case 1:
				nodes.Node1Post <- sendSlice
			case 2:
				nodes.Node2Post <- sendSlice
			case 3:
				nodes.Node3Post <- sendSlice
			}

		case "DELETE":
			key := fields[1]
			node := nodes.Node(key)

			switch node {
			case 0:
				nodes.Node0Delete <- key
			case 1:
				nodes.Node1Delete <- key
			case 2:
				nodes.Node2Delete <- key
			case 3:
				nodes.Node3Delete <- key
			}

		case "GET":
			key := fields[1]
			node := nodes.Node(key)

			switch node {
			case 0:
				io.WriteString(conn, fmt.Sprintf("%v: %v\n", key, nodes.Data0[key]))
			case 1:
				io.WriteString(conn, fmt.Sprintf("%v: %v\n", key, nodes.Data1[key]))
			case 2:
				io.WriteString(conn, fmt.Sprintf("%v: %v\n", key, nodes.Data2[key]))
			case 3:
				io.WriteString(conn, fmt.Sprintf("%v: %v\n", key, nodes.Data3[key]))
			}

		default:
			io.WriteString(conn, "Invalid Command "+fields[0]+"\n")
		}

	}

}
