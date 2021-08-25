package network

import (
	"fmt"
	"kv-store/store"
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
		input := strings.TrimSpace(string(buffer[:n]))
		fields := strings.Split(string(input), " ")

		if input == "STOP" {
			connection.WriteToUDP([]byte("Closing connection...\n"), addr)
			return
		}

		switch fields[0] {
		case "LIST":
			message := "\n"
			i := 0
			for _, v := range store.Data {
				i++
				message = message + strconv.Itoa(i) + ": " + v + "\n"
			}
			reply := []byte(message)
			connection.WriteToUDP(reply, addr)
		case "POST":
			var value string
			message := ""
			if len(fields) < 2 {
				message = "No value added!\n"
				connection.WriteToUDP([]byte(message), addr)
			}
			value = strings.Join(fields[1:], " ")
			store.PostChannel <- value
			message = fmt.Sprintf("%v added!\n", value)
			connection.WriteToUDP([]byte(message), addr)
		case "GET":
			requested, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Print(err)
			}
			connection.WriteToUDP([]byte(store.Data[requested]+"\n"), addr)
		case "DELETE":
			requested, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Print(err)
			}
			store.DelChannel <- requested
		}
		if err != nil {
			fmt.Print(err)
		}
	}
}
