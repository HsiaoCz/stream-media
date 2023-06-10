package storage

type RedisStorage interface{}

type Redis_Storage struct{}

func NewRedisStorage() *Redis_Storage {
	return &Redis_Storage{}
}
