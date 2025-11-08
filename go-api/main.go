package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/streadway/amqp"

    "southpark/adapters/rabbitmq"
    "southpark/app"
    "southpark/domain"
)

func main() {
    // RabbitMQ connection string - when running via Docker Compose service name will be "rabbitmq"
    rabbitURL := os.Getenv("RABBITMQ_URL")
    if rabbitURL == "" {
        rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
    }

    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    queueName := "southpark_messages"
    _, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    publisher := rabbitmq.NewRabbitMQPublisher(ch, queueName)
    msgService := app.NewMessageService(publisher)

    http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        var msg domain.Message
        if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        if err := msgService.SendMessage(msg); err != nil {
            if errors.Is(err, app.ErrInvalidMessage) {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusAccepted)
        fmt.Fprintln(w, "Message accepted")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Go API running on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}