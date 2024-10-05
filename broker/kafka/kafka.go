package kafka

import (
	"github.com/IBM/sarama"

	"github.com/thanhtranna/outbox"
)

// broker implements the MessageBroker interface
type broker struct {
	producer sarama.SyncProducer
}

// NewBroker constructor
func NewBroker(brokers []string, config *sarama.Config) (outbox.MessageBroker, error) {
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &broker{producer: producer}, nil
}

// Send delivers the message to kafka
func (b broker) Send(event outbox.Message) error {
	var headers []sarama.RecordHeader

	for k, v := range event.Headers {
		headers = append(headers, sarama.RecordHeader{
			Key:   sarama.ByteEncoder(k),
			Value: sarama.ByteEncoder(v),
		})
	}

	msg := &sarama.ProducerMessage{
		Topic:   event.Topic,
		Key:     sarama.StringEncoder(event.Key),
		Value:   sarama.StringEncoder(event.Body),
		Headers: headers,
	}
	_, _, err := b.producer.SendMessage(msg)

	return err
}
