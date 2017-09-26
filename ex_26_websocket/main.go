package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("websocket error")
			break
		}

		fmt.Println("client: " + reply)

		msg := "Received: " +reply
		fmt.Println("server: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("can not send")
			break
		}
	}
}


func main() {
	http.Handle("/", websocket.Handler(Echo))

	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		fmt.Println(err)
	}
}