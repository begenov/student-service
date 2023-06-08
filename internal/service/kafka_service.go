package service

import (
	"context"
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
				// Обработка прочитанных сообщений
				student, err := s.repo.GetStudentsByCoursesID(ctx, string(message.Value))
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
	<-ctx.Done()

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
