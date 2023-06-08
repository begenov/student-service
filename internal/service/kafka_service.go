package service

import (
	"context"
	"fmt"
	"log"

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

	log.Println("Message sent to Kafka:", message)
	return nil

}

func (s *KafkaService) Read(ctx context.Context) {
	partitions, err := s.consumer.Consumer.Partitions("students-request")
	if err != nil {
		log.Fatalln("Failed to get partitions:", err)
	}

	for _, partition := range partitions {
		pc, err := s.consumer.Consumer.ConsumePartition("students-request", partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln("Failed to start consumer for partition", partition, ":", err)
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for message := range pc.Messages() {
				res := ""
				for _, v := range message.Value {
					if v == '"' {
						continue
					}
					res += string(v)
				}
				// Обработка прочитанных сообщений
				student, err := s.repo.GetStudentsByCoursesID(ctx, res)
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println(student)
				if err := s.producer.SendMessage("students-response", student); err != nil {
					log.Println(err, "send message")
					return
				}

			}
		}(pc)
	}
	<-ctx.Done()

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
