package repository

type Repository interface {
	SendChatNotification(authorization string) error
}

type repository struct {
	//pubsub *pubsub.Client
}

func NewRepository() Repository {
	return &repository{}
}
