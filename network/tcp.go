package network

import (
	"bufio"
	"fmt"
	"io"
	"kv-store/store"
	"log"
	"net"
	"strconv"
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
			var value string
			if len(fields) < 2 {
				io.WriteString(conn, "No value added! \n")
			}
			fieldArr := strings.Split(fields[1], "_")
			if len(fieldArr) > 1 {
				value = strings.Join(fieldArr, " ")
			} else {
				value = fields[1]
			}
			store.PostChannel <- string(value)
			io.WriteString(conn, value+" added!\n")
		case "LIST":

			for i := 0; i < len(store.Data)+1; i++ {
				if store.Data[i] != "" {
					io.WriteString(conn, fmt.Sprint(i)+": "+store.Data[i]+"\n")
				}
			}

		case "SET":
			if len(fields) < 3 {
				io.WriteString(conn, "Format should be <int Key> <string Value> \n")
			}
			var chanSlice [2]string
			var value string
			keyInt, err := strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal("Fatal error")
			}
			chanSlice[0] = fields[1]
			fieldArr := strings.Split(fields[2], "_")
			if len(fieldArr) > 1 {
				value = strings.Join(fieldArr, " ")
			} else {
				value = fields[2]
			}
			if store.Data[keyInt] != "" {
				io.WriteString(conn, "Key is already in use. Try the LIST command to see keys in use \n")
			} else {
				chanSlice[1] = value
				store.SetChannel <- chanSlice
			}
		case "DELETE":
			keyInt, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Print(err)
			}
			store.DelChannel <- keyInt
		case "GET":
			requested, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Print(err)
			}
			io.WriteString(conn, store.Data[requested])
		case "STOP":
			io.WriteString(conn, "Closing connection...")
			conn.Close()
		default:
			io.WriteString(conn, "Invalid Command "+fields[0]+"\n")
		}

	}

}
