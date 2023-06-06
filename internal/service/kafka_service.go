package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"

	"github.com/begenov/student-service/internal/repository"
	"github.com/begenov/student-service/pkg/kafka"
)

type KafkaService struct {
	repo     repository.Students
	producer *kafka.Producer
	concumer *kafka.Consumer
}

func NewKafkaSerivce(repo repository.Students, producer *kafka.Producer, concumer *kafka.Consumer) *KafkaService {
	return &KafkaService{
		repo:     repo,
		producer: producer,
		concumer: concumer,
	}
}

func (s *KafkaService) Read(ctx context.Context) {
	for {
		responseHandler := func(message string) {
			// Добавьте здесь логику обработки ответа от Kafka

			students, err := s.repo.GetStudentsByCoursesID(ctx, message)
			if err != nil {

				log.Println("Failed to send students to course:", err)
				return
			}

			var buf bytes.Buffer
			encoder := gob.NewEncoder(&buf)
			err = encoder.Encode(students)
			if err != nil {
				log.Println("Failed to serialize student:", err)
				return
			}

			m := buf.Bytes()

			s.producer.SendMessage("students", string(m))

		}

		// Потребляем сообщения из Kafka
		err := s.concumer.ConsumeMessages("students", responseHandler)
		if err != nil {
			log.Println("Failed to consume messages from Kafka:", err)
			return
		}
	}
}

func (s *KafkaService) SendMessages(topic string, message string) error {
	return nil
}

func (s *KafkaService) ConsumeMessages(topic string, handler func(message string)) error {
	return nil
}

func (s *KafkaService) Close() {
	_ = s.concumer.Close()
	_ = s.producer.Close()
}
