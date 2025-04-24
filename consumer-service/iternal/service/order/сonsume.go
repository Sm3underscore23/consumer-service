package order

import (
	"consumer-service/iternal/model"
	"consumer-service/iternal/repository"
	"context"
	"log"

	"github.com/IBM/sarama"
)

type messageService struct {
	repo repository.OrderRepository
}

func NewMessageService(ctx context.Context, repo repository.OrderRepository, assignor, kafkaConsumerGroupstring string, brokerList []string, topicName string) error {
	consumerGroupHandler := messageService{repo: repo}

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
		return model.ErrCreateConsumerGroup
	}

	err = consumerGroup.Consume(ctx, []string{topicName}, &consumerGroupHandler)

	return err
}

func (consumer *messageService) Setup(sarama.ConsumerGroupSession) error {
	log.Println("consumer - setup")
	return nil
}

func (consumer *messageService) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("consumer - cleanup")
	return nil
}

func (consumer *messageService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := consumer.processOrder(session.Context(), message.Value)
		if err != nil {
			log.Printf("failed processing order: %s", err)
			return err
		}
		session.MarkMessage(message, "")
	}
	return nil
}
