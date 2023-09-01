package domain

type User struct {
	ID       int
	Login    string
	Password string
}

type UserRepository interface {
	Create(user *User) error
	GetByLogin(email string) error
	GetById(id string) error
}
