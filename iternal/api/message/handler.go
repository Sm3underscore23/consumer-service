package message

import "consumer-service/iternal/service"

type Implementation struct {
	messageService service.MessageService
}

func NewMessageImplementation(messageService service.MessageService) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
