package nodes

import "hash/fnv"

var (
	NumOfNodes uint64 = 4

	// Node 0
	Data0       = make(map[string]string)
	Node0Post   = make(chan []string)
	Node0Set    = make(chan []string)
	Node0Delete = make(chan string)

	//Node 1
	Data1       = make(map[string]string)
	Node1Post   = make(chan []string)
	Node1Set    = make(chan []string)
	Node1Delete = make(chan string)
	//Node 2
	Data2       = make(map[string]string)
	Node2Post   = make(chan []string)
	Node2Set    = make(chan []string)
	Node2Delete = make(chan string)
	//Node 3
	Data3       = make(map[string]string)
	Node3Post   = make(chan []string)
	Node3Set    = make(chan []string)
	Node3Delete = make(chan string)
)

//Hash Function

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64() % NumOfNodes
}

func Node0() {

	for {
		select {
		case postRequest := <-Node0Post:
			k := postRequest[0]
			v := postRequest[1]
			Data0[k] = v

		case setRequest := <-Node0Set:
			k := setRequest[0]
			v := setRequest[1]
			Data0[k] = v

		case delRequest := <-Node0Delete:
			k := delRequest[0]

			delete(Data0, string(k))

		}
	}

}
func Node1() {

	for {
		select {
		case postRequest := <-Node1Post:
			k := postRequest[0]
			v := postRequest[1]
			Data1[k] = v

		case setRequest := <-Node1Set:
			k := setRequest[0]
			v := setRequest[1]
			Data1[k] = v

		case delRequest := <-Node1Delete:
			k := delRequest[0]

			delete(Data1, string(k))

		}
	}

}

func Node2() {

	for {
		select {
		case postRequest := <-Node2Post:
			k := postRequest[0]
			v := postRequest[1]
			Data2[k] = v

		case setRequest := <-Node2Set:
			k := setRequest[0]
			v := setRequest[1]
			Data2[k] = v

		case delRequest := <-Node2Delete:
			k := delRequest[0]

			delete(Data2, string(k))

		}
	}

}
func Node3() {

	for {
		select {
		case postRequest := <-Node3Post:
			k := postRequest[0]
			v := postRequest[1]
			Data3[k] = v

		case setRequest := <-Node3Set:
			k := setRequest[0]
			v := setRequest[1]
			Data3[k] = v

		case delRequest := <-Node3Delete:
			k := delRequest[0]

			delete(Data3, string(k))

		}
	}

}
