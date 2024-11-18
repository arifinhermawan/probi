package rmq

import (
	"context"
	"encoding/json"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
)

func (svc *Service) PublishMessage(ctx context.Context, req PublishMessageReq) error {
	ctx, span := tracer.StartSpanFromContext(ctx, tracer.Service+"PublishMessage")
	defer span.End()

	bytes, err := json.Marshal(req.Message)
	if err != nil {
		log.Error(ctx, nil, err, "[PublishMessage] json.Marshal() got error")
		return err
	}

	err = svc.rmq.PublishMessage(ctx, rabbitmq.PublishMessageReq{
		Exchange:    req.Exchange,
		RouteKey:    req.RouteKey,
		IsMandatory: req.IsMandatory,
		IsImmediate: req.IsImmediate,
		Message:     bytes,
	})

	if err != nil {
		log.Error(ctx, nil, err, "[PublishMessage] svc.rmq.PublishMessage() got error")
		return err
	}

	return nil
}
