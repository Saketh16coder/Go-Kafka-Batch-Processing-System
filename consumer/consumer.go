package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"Demo/db"
	"Demo/model"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(workerID int) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "transactions",
		GroupID:  "txn-group", // 👈 Consumer Group
		MinBytes: 1,           // minimum fetch size
		MaxBytes: 10e6,        // maximum fetch size (10MB)
	})

	fmt.Printf("🚀 Worker %d started\n", workerID)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error reading message:", err)
			continue
		}

		var txn model.Transaction

		err = json.Unmarshal(msg.Value, &txn)
		if err != nil {
			log.Println("❌ JSON unmarshal error:", err)
			continue
		}

		fmt.Printf("📥 Worker %d received: %+v\n", workerID, txn)

		insertToDB(workerID, txn)
	}
}

func insertToDB(workerID int, txn model.Transaction) {

	query := `
		INSERT INTO transactions (user_id, amount, type)
		VALUES ($1, $2, $3)
	`

	_, err := db.DB.Exec(query, txn.UserID, txn.Amount, txn.Type)
	if err != nil {
		log.Println("❌ DB insert error:", err)
		return
	}

	fmt.Printf("✅ Worker %d inserted user:%d amount:%d type:%s\n",
		workerID, txn.UserID, txn.Amount, txn.Type)
}
