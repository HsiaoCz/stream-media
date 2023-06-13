package storage

type Storage struct {
	Ms *Mysql_Storage
	Rs *RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		Ms: NewMysqlStorage(),
		Rs: NewRedisStorage(),
	}
}
