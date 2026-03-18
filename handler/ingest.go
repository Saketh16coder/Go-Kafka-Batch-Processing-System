package handler

import (
	"encoding/json"
	"net/http"

	"Demo/model"
	"Demo/producer"
)

func IngestHandler(w http.ResponseWriter, r *http.Request) {

	var txn model.Transaction

	err := json.NewDecoder(r.Body).Decode(&txn)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = txn.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 🔥 Send to Kafka instead of DB
	err = producer.Publish(txn)
	if err != nil {
		http.Error(w, "Kafka error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sent to Kafka",
	})
}
