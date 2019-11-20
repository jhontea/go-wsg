package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func subscriber(w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer websocket.Close()

	for {
		var message MessageType
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			break
		}

		fmt.Printf("%+v\n", message)

		err = websocket.WriteJSON(message)
		if err != nil {
			log.Printf("error occurred while writing message to client: %v", err)
			break
		}
	}
}

func writeMessage(host string, message MessageType) {
	u := url.URL{Scheme: "ws", Host: host, Path: "/subscriber"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	err = conn.WriteJSON(message)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
