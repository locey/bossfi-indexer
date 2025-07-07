package mq

import (
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/log"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var KafkaProducer *kafka.Writer
var KafkaConsumer *kafka.Reader

func InitKafka() {
	// 初始化生产者
	KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(config.Conf.Kafka.Brokers...),
		Topic:    config.Conf.Kafka.Topic,
		Balancer: &kafka.Hash{},
		Async:    true,
	}

	// 初始化消费者
	KafkaConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.Conf.Kafka.Brokers,
		Topic:    config.Conf.Kafka.Topic,
		GroupID:  config.Conf.Kafka.GroupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	log.Logger.Info("Kafka initialized successfully")
}

func CloseKafka() {
	if KafkaProducer != nil {
		if err := KafkaProducer.Close(); err != nil {
			log.Logger.Error("Failed to close Kafka producer", zap.Error(err))
		}
	}
	if KafkaConsumer != nil {
		if err := KafkaConsumer.Close(); err != nil {
			log.Logger.Error("Failed to close Kafka consumer", zap.Error(err))
		}
	}
}
