package storage

type MySqlStorage interface {
	UserLogin() error
	UserRegister() error
	UserDelete() error
	GetUserByName() error
}

type RedisStorage interface{}
