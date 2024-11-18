package rabbitmq

type PublishMessageReq struct {
	Exchange    string
	RouteKey    string
	IsMandatory bool
	IsImmediate bool
	Message     []byte
}
