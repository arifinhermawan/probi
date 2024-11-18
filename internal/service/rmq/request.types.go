package rmq

type PublishMessageReq struct {
	Exchange    string
	RouteKey    string
	IsMandatory bool
	IsImmediate bool
	Message     interface{}
}
