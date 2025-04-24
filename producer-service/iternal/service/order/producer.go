package order

import (
	"event-generator/iternal/model"
	"event-generator/iternal/service"
	"time"

	"github.com/IBM/sarama"
)

type messageService struct {
	producer  sarama.SyncProducer
	topicName string
}

func NewKafkaService(brokerList []string, topicName string) (service.OrderSenderService, error) {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Retry.Backoff = time.Millisecond * 250
	config.Producer.Return.Successes = true
	_ = config.Producer.Partitioner
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, model.ErrStartProducer
	}
	return &messageService{producer: producer, topicName: topicName}, nil
}

func (s *messageService) Close() error {
	return s.producer.Close()
}
