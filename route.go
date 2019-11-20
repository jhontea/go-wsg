package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func route() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var req map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&req)
		defer r.Body.Close()

		result := executeQuery(req["query"].(string), schema)
		json.NewEncoder(w).Encode(result)
	})
	myRouter.HandleFunc("/subscriber", subscriber)

	log.Fatal(http.ListenAndServe(":1200", myRouter))
}
