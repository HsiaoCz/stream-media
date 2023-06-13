package storage

type RedisStorage struct {
	Rs userRedisI
}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{
		Rs: newUserRedis(),
	}
}

func (r *RedisStorage) initStore() error {
	return nil
}
