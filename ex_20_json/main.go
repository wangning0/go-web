// package main

// import (
// 	"fmt"
// 	"encoding/json"
// )

// type Server struct {
// 	ServerName string
// 	serverIp string
// }

// type ServerLists struct {
// 	Servers []Server
// }

// func main() {
// 	var s ServerLists
// 	str := `{"servers": [{"serverName": "shanghai", "serverIp": "127.0.0.1"}, {"serverName": "beijing", "serverIp": "127.0.0.1"}]}`
// 	json.Unmarshal([]byte(str), &s)
// 	fmt.Println(s);
// }


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
// 	var f interface{}

// 	err := json.Unmarshal(b, &f)
// 	if err != nil {
// 		// ...
// 	}
// 	fmt.Println(f)
// }


// 生成JSON

package main

import (
	"fmt"
	"encoding/json"
)

type Server struct {
	// 后缀的json:"" 是为了将输出的形式改成小写
	ServerName string `json:"serverName"`
	ServerIp string `json:"serverIp"`
}

type ServerLists struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s ServerLists
	s.Servers = append(s.Servers, Server{ServerName: "shanghai", ServerIp: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "beijing", ServerIp: "127.0.0.1"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(b))
}