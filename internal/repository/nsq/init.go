package nsq

type nsqPublishProvider interface {
	Publish(topic string, body []byte) error
}

type Repository struct {
	nsq nsqPublishProvider
}

func NewNSQRepo(nsq nsqPublishProvider) *Repository {
	return &Repository{
		nsq: nsq,
	}
}
