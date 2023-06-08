package service

import (
	"bytes"
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/begenov/student-service/internal/repository"
	"github.com/begenov/student-service/pkg/kafka"
)

type KafkaService struct {
	repo     repository.Students
	producer *kafka.Producer
	consumer *kafka.Consumer
}

func NewKafkaSerivce(repo repository.Students, producer *kafka.Producer, consumer *kafka.Consumer) *KafkaService {
	return &KafkaService{
		repo:     repo,
		producer: producer,
		consumer: consumer,
	}
}

func (s *KafkaService) SendMessages(topic string, message string) error {
	if err := s.producer.SendMessage(topic, message); err != nil {
		return err
	}

	return nil

}

func (s *KafkaService) Read(ctx context.Context) {
	partitions, err := s.consumer.Consumer.Partitions("students-request")
	if err != nil {
		log.Fatalln("Failed to get partitions:", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	for _, partition := range partitions {
		pc, err := s.consumer.Consumer.ConsumePartition("students-request", partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln("Failed to start consumer for partition", partition, ":", err)
		}

		go func(pc sarama.PartitionConsumer) {
			defer func() {
				pc.Close()
				wg.Done()
			}()

			for message := range pc.Messages() {
				res := getStringWithoutQuotes(message.Value)

				student, err := s.repo.GetStudentsByCoursesID(ctx, res)
				if err != nil {
					log.Println(err)
					return
				}

				if err := s.producer.SendMessage("students-response", student); err != nil {
					log.Println(err, "send message")
					return
				}
			}
		}(pc)
	}

	wg.Wait()
}

func getStringWithoutQuotes(input []byte) string {
	var buffer bytes.Buffer

	for _, v := range input {
		if v == '"' {
			continue
		}
		buffer.WriteByte(v)
	}

	return buffer.String()
}

func (s *KafkaService) ConsumeMessages(topic string, handler func(message string)) error {
	err := s.consumer.ConsumeMessages(topic, handler)
	if err != nil {
		log.Println("Failed to consume messages from Kafka:", err)
		return err
	}

	return nil
}

func (s *KafkaService) Close() {
	if err := s.consumer.Close(); err != nil {
		log.Println(err)
		return
	}
	if err := s.producer.Close(); err != nil {
		log.Println(err)
		return
	}
}
