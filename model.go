package main

// Message - Our struct for all messages
type Message struct {
	Message string `json:"message"`
}

// Messages let's declare a global Messages array
// that we can then populate in our main function
// to simulate a database
var Messages []Message
