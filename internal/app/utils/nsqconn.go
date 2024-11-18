package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/arifinhermawan/blib/log"
	"github.com/nsqio/go-nsq"
)

func InitNSQProducer(ctx context.Context, nsqd string) (*nsq.Producer, error) {
	producer, err := nsq.NewProducer(nsqd, nsq.NewConfig())
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func GracefulShutDownProducer(ctx context.Context, publisher *nsq.Producer) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	<-stopCh
	log.Info(ctx, nil, nil, "Received shutdown signal, stopping producer...")
	publisher.Stop()
}
