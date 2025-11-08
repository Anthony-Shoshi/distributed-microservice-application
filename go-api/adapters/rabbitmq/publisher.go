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
    body, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    return r.Channel.Publish(
        "",          // exchange
        r.Queue,     // routing key = queue name
        false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}