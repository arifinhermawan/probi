package reminder

import (
	"context"
	"fmt"

	"github.com/arifinhermawan/blib/tracer"
	"github.com/nsqio/go-nsq"
)

func (h *Handler) SendReminderConsumer(ctx context.Context, message *nsq.Message) {
	ctx, txn := tracer.StartTransaction(ctx, tracer.Handler+"SendReminderConsumer", tracer.MQTransaction)
	defer txn.End()

	fmt.Println("MESSAGE:", string(message.Body))
}
