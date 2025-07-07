package consumer

import (
	"bossfi-indexer/src/app/model"
	"bossfi-indexer/src/core/log"
	"bossfi-indexer/src/core/mq"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"sync"
	"time"
)

func StartKafkaConsumer() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			msg, err := mq.KafkaConsumer.ReadMessage(context.Background())
			if err != nil {
				log.Logger.Error("Failed to read Kafka message", zap.Error(err))
				time.Sleep(5 * time.Second) // 避免频繁重试
				continue
			}

			var event model.TokenEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Logger.Error("Failed to unmarshal Kafka message",
					zap.Error(err),
					zap.ByteString("value", msg.Value))
				continue
			}

			// 处理消息
			if err := ProcessEvent(&event); err != nil {
				log.Logger.Error("Failed to process event",
					zap.Error(err),
					zap.Any("event", event))
			}
		}
	}()

	wg.Wait()
}

func ProcessEvent(event *model.TokenEvent) error {
	// 根据业务需求处理事件
	// 例如: 更新缓存、通知其他服务等

	return nil
}
