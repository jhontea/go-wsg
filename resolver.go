package main

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"messages": &graphql.Field{
			Type:        graphql.NewList(messageType),
			Description: "List of messages",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return MessageTypes, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createMessage": &graphql.Field{
			Type:        messageType, // the return type for this field
			Description: "Create new message",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				text, _ := params.Args["text"].(string)

				newID := RandStringRunes(8)

				newMessage := MessageType{
					ID:   newID,
					Text: text,
				}

				MessageTypes = append(MessageTypes, newMessage)
				writeMessage("localhost:1200", newMessage)

				return newMessage, nil
			},
		},
	},
})
