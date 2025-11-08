package ports

import "southpark/domain"

// MessagePublisher is the port for publishing messages
type MessagePublisher interface {
    Publish(msg domain.Message) error
}