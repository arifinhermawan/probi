package reminder

import (
	"context"
	"fmt"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/arifinhermawan/probi/internal/handler/mq"
	"github.com/arifinhermawan/probi/internal/lib/constants"
	"github.com/streadway/amqp"
)

func (h *Handler) SendReminderEmail(ctx context.Context) {
	ctx, txn := tracer.StartTransaction(ctx, tracer.Handler+"SendReminderEmail", constants.ServiceNameMQ)
	defer txn.End()

	cfg := h.lib.GetConfig().Channel.Queue
	messages, err := h.rmq.Consume(
		cfg.ReminderSendEmail.Name,
		cfg.ReminderSendEmail.Consumer,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error(ctx, nil, err, "[SendReminderEmail] h.rmq.Consume() got error")
		return
	}

	mq.ConsumeMessages(ctx, messages, func(msg amqp.Delivery) error {
		fmt.Println("this is the message", string(msg.Body))
		return nil
	})
}
