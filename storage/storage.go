package storage

type Storage struct {
	MySqlStorage
	RedisStorage
}

func NewStorage(mysql_Storage MySqlStorage, redis_Storage RedisStorage) *Storage {
	return &Storage{
		MySqlStorage: mysql_Storage,
		RedisStorage: redis_Storage,
	}
}
