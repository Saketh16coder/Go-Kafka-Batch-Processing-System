package main

import (
	"log"
	"net/http"

	"Demo/db"
	"Demo/handler"
	"Demo/producer"
)

func main() {
	db.Connect()
	producer.InitProducer()

	http.HandleFunc("/ingest", handler.IngestHandler)

	log.Println("🚀 Producer API running on :8080")
	http.ListenAndServe(":8080", nil)
}
