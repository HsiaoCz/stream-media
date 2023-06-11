package storage

type Storage struct {
	MySqlStorage
	RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		MySqlStorage: NewMysqlStorage().um,
		RedisStorage: NewRedisStorage(),
	}
}
