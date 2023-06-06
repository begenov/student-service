package kafka

import (
	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
	}, nil
}

func (c *Consumer) ConsumeMessages(topic string, handler func(message string)) error {
	partitions, err := c.consumer.Partitions(topic)
	if err != nil {

		return err
	}

	for _, partition := range partitions {
		pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {

			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				handler(string(message.Value))
			}
		}(pc)
	}

	return nil
}

func (c *Consumer) Close() error {
	if c.consumer != nil {
		return c.consumer.Close()
	}
	return nil
}
