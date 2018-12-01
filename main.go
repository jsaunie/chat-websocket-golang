package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	hub := make([]*websocket.Conn, 0)

	http.Handle("/chat/", websocket.Handler(func(ws *websocket.Conn) {
		data := map[string]interface{}{}
		hub = append(hub, ws)

		for {
			err := websocket.JSON.Receive(ws, &data)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(data)
			for i := range hub {
				err2 := websocket.JSON.Send(hub[i], data)
				if err2 != nil {
					fmt.Println(err2)
					break
				}
			}
		}
		err3 := ws.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}))
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println(err)
	}
}
