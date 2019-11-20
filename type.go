package main

import (
	"github.com/graphql-go/graphql"
)

// MessageType - Our struct for all messages
type MessageType struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// MessageTypes let's declare a global Messages array
// that we can then populate in our main function
// to simulate a database
var MessageTypes []MessageType

var messageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Message",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
	},
})
