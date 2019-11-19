package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/health", healthCheckHandler)

	myRouter.HandleFunc("/messages", getMessages)
	myRouter.HandleFunc("/message", createMessage).Methods("POST")

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	myRouter.HandleFunc("/subscriber", subscriber)

	log.Fatal(http.ListenAndServe(":1200", myRouter))
}
