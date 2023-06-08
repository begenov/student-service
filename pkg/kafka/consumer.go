package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	Consumer sarama.Consumer
	done     chan struct{}
}

func NewConsumer(brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()

	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Consumer: consumer,
		done:     make(chan struct{}),
	}, nil
}

func (c *Consumer) ConsumeMessages(topic string, handler func(message string)) error {
	partitions, err := c.Consumer.Partitions(topic)
	if err != nil {
		log.Println("Failed to retrieve partitions:", err)
		return err
	}

	if len(partitions) == 0 {
		log.Println("No partitions found for the topic:", topic)
		return nil
	}

	for _, partition := range partitions {
		pc, err := c.Consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Println("Failed to start consumer for partition", partition, ":", err)
			return err
		}
		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for message := range pc.Messages() {

				handler(string(message.Value))
			}
		}(pc)
	}

	return nil

}

func (c *Consumer) Stop() {
	close(c.done)
}

func (c *Consumer) Close() error {
	c.Stop()

	if c.Consumer != nil {
		return c.Consumer.Close()
	}
	return nil
}
