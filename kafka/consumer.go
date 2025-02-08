package kafka

import (
	"HyperLightLogistics-Go/internal/service"
	"context"
	"log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.ConsumerGroup
	topic    string
	service  *service.DeliveryService
}

type consumerGroupHandler struct {
	service *service.DeliveryService
}

func NewKafkaConsumer(brokers []string, groupID, topic string, service *service.DeliveryService) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumerGroup,
		topic:    topic,
		service:  service,
	}, nil
}

func (k *KafkaConsumer) StartConsuming(ctx context.Context) {
	handler := consumerGroupHandler{k.service}

	for {
		if err := k.consumer.Consume(ctx, []string{k.topic}, handler); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

func (h consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received Kafka message: %s", string(msg.Value))

		err := h.service.DeliveryInitialization(msg.Value)
		if err != nil {
			log.Printf("Failed to start delivery: %v", err)
		}

		session.MarkMessage(msg, "")
	}
	return nil
}
