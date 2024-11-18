package nsq

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
)

func (r *Repository) PublishMessageToNSQ(ctx context.Context, topic string, message []byte) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.MQ+"PublishMessageToNSQ")
	defer span.End()

	err := r.nsq.Publish(topic, message)
	if err != nil {
		log.Error(ctx, map[string]interface{}{
			"topic":   topic,
			"message": string(message),
		}, err, "[PublishMessageToNSQ] r.nsq.Publish()")

		return err
	}

	return nil
}
