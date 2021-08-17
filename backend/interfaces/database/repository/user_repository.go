package repository

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/handler"
)

type UserRepository struct {
	handler.SqlHandler
}

func (repo *UserRepository) Insert(u domain.User) (id int, err error) {
	result, err := repo.NamedExec(`
		INSERT INTO users (
			name,
			password
		) VALUES (
			:name,
			:password
		)
	`, u)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id = int(id64)
	return id, err
}
