package network

import (
	"fmt"
	"kv-store/nodes"
	"net/http"
	"strings"
)

func HttpHandleFuncs() {

	http.HandleFunc("/get/", GetData)
	http.HandleFunc("/post/", Post)
	http.HandleFunc("/delete/", Delete)
	http.HandleFunc("/set/", SetData)
	http.HandleFunc("/list", ListTest)
	http.ListenAndServe(":8001", nil)

}

//HTTP HandleFuncs

func GetData(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	key := parts[2]
	node := nodes.Node(key)

	switch node {
	case 0:
		fmt.Fprintf(w, nodes.Data0[key])
	case 1:
		fmt.Fprintf(w, nodes.Data1[key])
	case 2:
		fmt.Fprintf(w, nodes.Data2[key])
	case 3:
		fmt.Fprintf(w, nodes.Data3[key])

	}
	fmt.Println(key)
	fmt.Println(node)
}

func Post(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var postSlice = make([]string, 2, 2)
	fields := strings.Split(parts[2], " ")
	key := fields[0]
	valueSlice := strings.Split(parts[3], "%20")
	value := strings.Join(valueSlice, " ")
	node := nodes.Node(key)

	postSlice[0] = key
	postSlice[1] = value

	switch node {
	case 0:
		nodes.Node0Post <- postSlice
	case 1:
		nodes.Node1Post <- postSlice
	case 2:
		nodes.Node2Post <- postSlice
	case 3:
		nodes.Node3Post <- postSlice
	}
	fmt.Println(node)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := parts[2]
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
}

func SetData(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
	}
	fields := strings.Split(parts[2], " ")

	var sendData = make([]string, 2, 2)
	key := fields[0]
	node := nodes.Node(key)
	value := strings.Join(fields[1:], " ")
	sendData[0] = key
	sendData[1] = value

	switch node {
	case 0:
		nodes.Node0Set <- sendData
	}
}

func ListTest(w http.ResponseWriter, r *http.Request) {
	for _, v := range nodes.Data3 {
		fmt.Fprintf(w, fmt.Sprintf("value from Data 3: %v\n", v))
	}
	for _, v := range nodes.Data2 {
		fmt.Fprintf(w, fmt.Sprintf("value from Data 2: %v\n", v))
	}
}
