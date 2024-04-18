package kafka

import (
	configpkg "evrone_api_gateway/internal/pkg/config"
	otlp_pkg "evrone_api_gateway/internal/pkg/otlp"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type producer struct {
	logger                       *zap.Logger
	investmentPaymentTransaction *kafka.Writer
}

func NewProducer(config *configpkg.Config, logger *zap.Logger) *producer {
	return &producer{
		logger: logger,
		investmentPaymentTransaction: &kafka.Writer{
			Addr:                   kafka.TCP(config.Kafka.Address...),
			Topic:                  config.Kafka.Topic.InvestmentPaymentTransaction,
			Balancer:               &kafka.Hash{},
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
			Async:                  true,
			Completion: func(messages []kafka.Message, err error) {
				if err != nil {
					logger.Error("kafka investmentCreated", zap.Error(err))
				}
				for _, message := range messages {
					logger.Sugar().Info(
						"kafka investmentCreated message",
						zap.Int("partition", message.Partition),
						zap.Int64("offset", message.Offset),
						zap.String("key", string(message.Key)),
						zap.String("value", string(message.Value)),
					)
				}
			},
		},
	}
}

func (p *producer) buildMessageWithTracing(key string, value []byte, otlpSpan otlp_pkg.Span) kafka.Message {
	return kafka.Message{
		Key:   []byte(key),
		Value: value,
		Headers: []kafka.Header{
			{
				Key:   "trace_id",
				Value: []byte(otlpSpan.SpanContext().TraceID().String()),
			},
			{
				Key:   "span_id",
				Value: []byte(otlpSpan.SpanContext().SpanID().String()),
			},
		},
	}
}

func (p *producer) Close() {
	if err := p.investmentPaymentTransaction.Close(); err != nil {
		p.logger.Error("error during close writer investmentCreated", zap.Error(err))
	}
}
