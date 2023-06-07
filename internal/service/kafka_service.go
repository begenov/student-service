package service

import (
	"context"
	"fmt"
	"log"

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
	for {

		t, err := s.consumer.Consumer.Partitions("students")
		fmt.Println("----", t, err)
		if err != nil {
			log.Println(err)
			return
		}

		for _, v := range t {

			pc, err := s.consumer.Consumer.ConsumePartition("students", v, -2)
			if err != nil {
				log.Println(err, "---")

				return
			}
			fmt.Printf("pc: %v\n", pc)
		}

	}
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
