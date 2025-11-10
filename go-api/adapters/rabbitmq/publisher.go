package rabbitmq

import (
    "encoding/json"

    "github.com/streadway/amqp"
    "southpark/domain"
    "southpark/ports"
)

type RabbitMQPublisher struct {
    Channel *amqp.Channel
    Queue   string
}

func NewRabbitMQPublisher(ch *amqp.Channel, queue string) ports.MessagePublisher {
    return &RabbitMQPublisher{
        Channel: ch,
        Queue:   queue,
    }
}

func (r *RabbitMQPublisher) Publish(msg domain.Message) error {

    _, err := r.Channel.QueueDeclare(
        r.Queue, // queue name
        true,    // durable
        false,   // auto-delete
        false,   // exclusive
        false,   // no-wait
        nil,     // args
    )
    if err != nil {
        return err
    }

    body, err := json.Marshal(msg)
    if err != nil {
        return err
    }

    return r.Channel.Publish(
        "",
        r.Queue,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}