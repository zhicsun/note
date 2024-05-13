package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"go.opentelemetry.io/otel"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

var (
	brokers = []string{
		"web-kafka-course-01.db.sit13.dom:9092",
		"web-kafka-course-02.db.sit13.dom:9092",
		"web-kafka-course-03.db.sit13.dom:9092",
	}
	topic       = "dev"
	clientId    = "client"
	group       = "group"
	kafkaConfig = sarama.NewConfig()
)

func TestKafkaConn(t *testing.T) {
	ctx := context.Background()
	GetKafkaSyncSend(ctx, brokers, kafkaConfig, topic)
	GetKafkaAsyncSend(ctx, brokers, kafkaConfig, topic)
}

func TestKafkaConsumer(t *testing.T) {
	KafkaConsumer(context.Background())
}

func KafkaConsumer(ctx context.Context) {
	//sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	kafkaConfig.ClientID = clientId
	kafkaConfig.Version = sarama.DefaultVersion
	kafkaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer := &Consumer{
		ready: make(chan bool),
	}

	handle := otelsarama.WrapConsumerGroupHandler(consumer)
	ctx, cancel := context.WithCancel(ctx)
	client, err := sarama.NewConsumerGroup(brokers, group, kafkaConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{topic}, handle); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(err.Error())
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()
	fmt.Println("启动成功", <-consumer.ready)

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true
	consumptionIsPaused := false
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		fmt.Println("Resuming consumption")
	} else {
		client.PauseAll()
		fmt.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Printf("message channel was closed")
				return nil
			}

			fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func setKafkaConfig() {
	//sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	kafkaConfig.ClientID = clientId
	kafkaConfig.Version = sarama.DefaultVersion
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy
	kafkaConfig.Producer.Flush.Frequency = 500 * time.Millisecond
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true
	kafkaConfig.Producer.Retry.Max = 10
	kafkaConfig.Producer.Return.Successes = true
}

func GetKafkaSyncSend(ctx context.Context, brokerList []string, config *sarama.Config, topic string) {
	setKafkaConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	producer = otelsarama.WrapSyncProducer(config, producer)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(time.Now().String()),
	}
	otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(msg))
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(partition, offset)
	fmt.Println(producer.Close())
}

func GetKafkaAsyncSend(ctx context.Context, brokerList []string, config *sarama.Config, topic string) {
	setKafkaConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		fmt.Println("Failed to start Sarama producer:", err)
	}

	producer = otelsarama.WrapAsyncProducer(config, producer)
	nowStr := time.Now().String()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(nowStr),
		Value: sarama.StringEncoder(nowStr),
	}
	otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(msg))
	producer.Input() <- msg
	fmt.Println(<-producer.Successes())

	go func() {
		for err := range producer.Errors() {
			fmt.Println("Failed to write access log entry:", err)
		}
	}()
	producer.AsyncClose()
}
