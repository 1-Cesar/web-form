package main

import (
	"log"
	"net/http"
	"web-form/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
