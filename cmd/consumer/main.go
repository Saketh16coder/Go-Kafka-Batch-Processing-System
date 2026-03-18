package main

import (
	"fmt"

	"Demo/consumer"
	"Demo/db"
)

func main() {

	fmt.Println("🚀 Starting Consumer Service...")

	// Connect DB
	db.Connect()

	// 🔥 Start multiple workers (consumer group)
	totalWorkers := 3

	for i := 1; i <= totalWorkers; i++ {
		go consumer.StartConsumer(i)
	}

	// Block forever
	select {}
}
