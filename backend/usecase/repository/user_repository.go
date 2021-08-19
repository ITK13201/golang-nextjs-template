package repository

import "github.com/ITK13201/golang-nextjs-template/backend/domain"

type UserRepository interface {
	Store(user domain.User) (int, error)
	FindById(id int) (*domain.User, error)
	FindAll() (domain.Users, error)
}
