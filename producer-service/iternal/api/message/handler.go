package message

import "event-generator/iternal/service"

type Implementation struct {
	messageService service.OrderSenderService
}

func NewMessageImplementation(messageService service.OrderSenderService) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
