package order

import (
	"consumer-service/iternal/repository"
	"consumer-service/iternal/service"
	"context"
	"log"

	"github.com/IBM/sarama"
)

type messageService struct {
	repo repository.OrderRepository
}
type Consumer struct {
}

func NewMessageService(ctx context.Context, repo repository.OrderRepository, assignor, kafkaConsumerGroupstring string, brokerList []string, topicName string) service.MessageService {
	consumerGroupHandler := Consumer{}

	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	switch assignor {
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "round-robin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokerList, kafkaConsumerGroupstring, config)
	if err != nil {
		log.Fatalf("failed to create consumer group: %s", err)
	}

	err = consumerGroup.Consume(ctx, []string{topicName}, &consumerGroupHandler)
	if err != nil {
		log.Fatalf("failed to consume via handler: %s", err)
	}

	return messageService{repo: repo}
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("consumer - setup")
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("consumer - cleanup")
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		session.MarkMessage(message, "geted")
	}
	return nil
}
