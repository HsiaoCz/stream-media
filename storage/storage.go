package storage

type Storage interface {
	MySqlStorage
	RedisStorage
}
