package storage

type userRedisI interface {
	GetUser(int) error
}

type userRedis struct{}

func newUserRedis() *userRedis {
	return &userRedis{}
}

func (u *userRedis) GetUser(userID int) error {
	return nil
}
