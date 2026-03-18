# 🚀 Go Kafka Batch Processing System

This project demonstrates a **production-style event-driven backend system** built using **Go, Kafka, and PostgreSQL**.

It simulates high-load traffic using batch requests and processes them asynchronously using Kafka consumers and worker pools.

---

# 🏗️ Architecture

```
Batch Client
     ↓
Go API (Producer)
     ↓
Kafka Topic (transactions)
     ↓
Consumer Group (Workers)
     ↓
PostgreSQL
```

---

# 📦 Features

* ✅ Batch processing (3 batches × 100 concurrent requests)
* ✅ Dynamic payload generation (random user_id, amount, type)
* ✅ Kafka-based decoupled architecture
* ✅ Consumer groups with parallel workers
* ✅ PostgreSQL data persistence
* ✅ Clean modular Go project structure

---

# 📁 Project Structure

```
Demo/
│
├── client/                 # Batch load generator
│   └── batch_client.go
│
├── handler/               # API layer
│   └── ingest.go
│
├── model/                 # Data models & validation
│   ├── transaction.go
│   └── validation.go
│
├── producer/              # Kafka producer
│   └── kafka.go
│
├── consumer/              # Consumer logic
│   └── consumer.go
│
├── cmd/
│   └── consumer/          # Consumer entry point
│       └── main.go
│
├── db/                    # PostgreSQL connection
│   └── db.go
│
├── main.go                # Producer API entry point
├── docker-compose.yml     # Kafka setup
├── go.mod
└── go.sum
```

---

# ⚙️ Setup Instructions

## 1️⃣ Start Kafka (Docker)

```bash
docker-compose up -d
```

---

## 2️⃣ Create Kafka Topic

```bash
docker exec -it <kafka_container> kafka-topics \
--create \
--topic transactions \
--bootstrap-server localhost:9092 \
--partitions 3 \
--replication-factor 1
```

---

## 3️⃣ Setup PostgreSQL

Create database:

```sql
CREATE DATABASE loadtesting;
```

Connect:

```sql
\c loadtesting
```

Create table:

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT NOT NULL,
    type VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

# 🚀 Running the System

## ▶️ Start Consumer (Workers)

```bash
go run cmd/consumer/main.go
```

---

## ▶️ Start Producer API

```bash
go run main.go
```

---

## ▶️ Run Batch Client

```bash
go run client/batch_client.go
```

---

# 📊 Expected Flow

* Client sends 300 requests (3 batches × 100)
* API publishes messages to Kafka
* Kafka distributes messages across consumers
* Workers process messages and store them in PostgreSQL

---

# 🧪 Verify Data

```sql
SELECT COUNT(*) FROM transactions;
```

Expected: ~300 rows

---

# 🔥 Key Concepts Covered

* Goroutines & concurrency in Go
* Kafka producer & consumer groups
* Worker pool pattern
* Event-driven architecture
* Database integration
* Batch processing

---

# ⚠️ Common Issues

| Issue                           | Fix                             |
| ------------------------------- | ------------------------------- |
| No data in DB                   | Ensure consumer is running      |
| Consumer not receiving messages | Check topic name                |
| DB connection error             | Verify port & credentials       |
| Kafka not working               | Ensure Docker containers are up |

---

# 🚀 Future Improvements

* Retry mechanism for failed messages
* Dead Letter Queue (DLQ)
* Rate limiting & backpressure
* Idempotency handling
* Structured logging (Zap/Logrus)
* Metrics & monitoring (Prometheus/Grafana)

---

# 👨‍💻 Tech Stack

* Go (Golang)
* Kafka (segmentio/kafka-go)
* PostgreSQL
* Docker

---

# 🎯 Summary

This project demonstrates how to move from a **simple synchronous API** to a **scalable, event-driven system** using Kafka and Go.

---
