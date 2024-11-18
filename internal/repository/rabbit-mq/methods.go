package rabbitmq

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	"github.com/streadway/amqp"
)

func (r *RMQRepo) PublishMessage(ctx context.Context, req PublishMessageReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.MQ+"PublishMessage")
	defer span.End()

	err := r.rmq.Publish(req.Exchange, req.RouteKey, req.IsMandatory, req.IsImmediate, amqp.Publishing{
		ContentType: "application/json",
		Body:        req.Message,
	})
	if err != nil {
		log.Error(ctx, nil, err, "[PublishMessage] r.rmq.Publish() got error")
		return err
	}

	return nil
}
