package repository

import "github.com/ITK13201/golang-nextjs-template/backend/domain"

type UserRepository interface {
	Insert(user domain.User) (int, error)
}
