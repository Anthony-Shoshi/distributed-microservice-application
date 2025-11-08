package app

import (
    "errors"

    "southpark/domain"
    "southpark/ports"
)

var ErrInvalidMessage = errors.New("author and body must not be empty")

type MessageService struct {
    Publisher ports.MessagePublisher
}	

func NewMessageService(publisher ports.MessagePublisher) *MessageService {
    return &MessageService{Publisher: publisher}
}

func (s *MessageService) SendMessage(msg domain.Message) error {
    if msg.Author == "" || msg.Body == "" {
        return ErrInvalidMessage
    }
    return s.Publisher.Publish(msg)
}