package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"reflect"
)

func main() {
	http.Handle("/ws/", websocket.Handler(connHandler))
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func connHandler(ws *websocket.Conn) {
	data := map[string]interface{}{}

	for {
		err := websocket.JSON.Receive(ws, &data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Server receive : %v\n", data)
		fmt.Printf(
			"Types received : %s\n",
			reflect.TypeOf(data["text"]),
		)
		err2 := websocket.JSON.Send(ws, data)
		if err2 != nil {
			fmt.Println(err2)
			break
		}
	}
	err3 := ws.Close()
	if err3 != nil {
		fmt.Println(err3)
	}
}
