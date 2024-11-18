package router

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/nsq/server"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/nsqio/go-nsq"
)

func RegisterConsumer(ctx context.Context, handler *server.Handlers, cfg *configuration.AppConfig) {
	consumers := []struct {
		channel  configuration.ChannelConfig
		handlers nsq.HandlerFunc
	}{
		{
			cfg.Channel.Reminder,
			nsq.HandlerFunc(func(message *nsq.Message) error {
				handler.Reminder.SendReminderConsumer(ctx, message)
				return nil
			}),
		},
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	wg := new(sync.WaitGroup)
	wg.Add(len(consumers))

	for _, con := range consumers {
		go func(consumerConfig configuration.ChannelConfig, consumerHandler nsq.HandlerFunc) {
			defer wg.Done()

			config := createConsumerConfig(consumerConfig.MaxAttempt, consumerConfig.MaxInFlight)
			consumer, err := nsq.NewConsumer(consumerConfig.Topic, consumerConfig.Channel, config)
			if err != nil {
				log.Fatal(ctx, nil, err, "[RegisterConsumer] failed to create consumer")
				return
			}

			consumer.AddHandler(consumerHandler)
			err = consumer.ConnectToNSQLookupds(cfg.NSQ.Lookupd)
			if err != nil {
				log.Fatal(ctx, nil, err, "[RegisterConsumer] failed to connect to NSQLookupds")
				return
			}

			<-stopCh
			log.Info(ctx, nil, nil, "[RegisterConsumer] Received shutdown signal, stopping consumer...")
			consumer.Stop()
		}(con.channel, con.handlers)
	}

	<-stopCh
	wg.Wait()
}

func createConsumerConfig(maxAttempts int, maxInFlight int) *nsq.Config {
	config := nsq.NewConfig()
	config.MaxAttempts = uint16(maxAttempts)
	config.MaxInFlight = maxInFlight

	return config
}
