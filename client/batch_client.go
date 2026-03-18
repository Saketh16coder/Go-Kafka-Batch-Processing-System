package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func makeRequest(batchID, requestID int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "http://localhost:8080/ingest"

	userID := requestID + (batchID * 1000)

	amount := rand.Intn(1000) + 1

	types := []string{"credit", "debit"}
	txnType := types[rand.Intn(len(types))]

	payloadMap := map[string]interface{}{
		"user_id": userID,
		"amount":  amount,
		"type":    txnType,
	}

	payloadBytes, _ := json.Marshal(payloadMap)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Batch %d Req %d error: %v\n", batchID, requestID, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Batch %d Req %d → user:%d amount:%d type:%s\n",
		batchID, requestID, userID, amount, txnType)
}
func runBatch(batchID int, total int) {
	fmt.Printf("\n Starting Batch %d\n", batchID)

	var wg sync.WaitGroup

	for i := 1; i <= total; i++ {
		wg.Add(1)
		go makeRequest(batchID, i, &wg)
	}

	wg.Wait()

	fmt.Printf("Batch %d completed\n", batchID)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	totalBatches := 3
	requestsPerBatch := 100

	start := time.Now()

	for i := 1; i <= totalBatches; i++ {
		runBatch(i, requestsPerBatch)
	}

	fmt.Println("\nAll batches done")
	fmt.Println("Total Time:", time.Since(start))
}
