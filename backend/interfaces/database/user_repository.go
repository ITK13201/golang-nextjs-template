package database

import {
	"github.com/ITK13201/go-next-test/backend/domain"

	"github.com/ITK13201/go-next-test/backend/infrastracture"
}

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	res, err := repo.Execute()
}