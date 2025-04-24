package order

import (
	"context"
	"encoding/json"
	"event-generator/iternal/model"
	"time"

	"github.com/IBM/sarama"
)

func (s *messageService) SendOrderData(ctx context.Context, orderRq model.OrderRequest) error {
	order := model.Order{
		OrderInfo:  orderRq,
		UpdateTime: time.Now(),
	}
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: s.topicName,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = s.producer.SendMessage(msg)
	return err
}
