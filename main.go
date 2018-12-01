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
			"Types received : %s - %s - %s\n",
			reflect.TypeOf(data["text"]),
			reflect.TypeOf(data["nb"]),
			reflect.TypeOf(data["bool"]),
		)
		data["text"] = "Bien reçu"
		data["nb"] = 789
		data["bool"] = false

		err2 := websocket.JSON.Send(ws, data)
		if err2 != nil {
			fmt.Println(err2)
			break
		}
	}
	ws.Close()
}
