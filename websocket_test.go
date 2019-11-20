package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// TestSubscriber test websocket subscribe
func TestSubscriber(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(subscriber))
	defer server.Close()

	u := "ws" + strings.TrimPrefix(server.URL, "http")
	socket, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer socket.Close()

	m := MessageType{
		ID:   "a",
		Text: "hello",
	}
	if err := socket.WriteJSON(&m); err != nil {
		t.Fatalf("%v", err)
	}

	var message MessageType
	err = socket.ReadJSON(&message)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := "hello"

	assert.Equal(t, expected, message.Text, "they should be equal")
}
