package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Response object
type Response struct {
	Data   interface{} `json:"data"`
	Remark string      `json:"remark"`
	Error  string      `json:"error,omitempty" `
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data:   Messages,
		Remark: "success retrieve data",
	}

	json.NewEncoder(w).Encode(response)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	var response Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// request payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Invalid request payload"
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()

	// validate message
	if (message == Message{}) {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Message cannot be empty"
		json.NewEncoder(w).Encode(response)
		return
	}

	Messages = append(Messages, message)

	writeMessage(r.Host, message)

	response.Data = message
	response.Remark = "success store message"

	json.NewEncoder(w).Encode(response)
}

func subscriber(w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer websocket.Close()

	for {
		var message Message
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

func writeMessage(host string, message Message) {
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
