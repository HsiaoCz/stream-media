package storage

type MySqlStorage interface {
	UserLogin() error
	UserRegister() error
	UserDelete() error
	GetUserByName() error
}
type Mysql_Storage struct{}

func NewMysqlStorage() *Mysql_Storage {
	return &Mysql_Storage{}
}

func (s *Mysql_Storage) UserRegister() error  { return nil }
func (s *Mysql_Storage) UserLogin() error     { return nil }
func (s *Mysql_Storage) GetUserByName() error { return nil }
func (s *Mysql_Storage) UserDelete() error    { return nil }
