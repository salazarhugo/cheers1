package repository

type Repository interface {
	SendChatNotification(userWithToken map[string][]string) error
}

type repository struct {
	//pubsub *pubsub.Client
}

func NewRepository() Repository {
	return &repository{}
}
