package repository

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/handler"
)

type UserRepository struct {
	handler.SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (int, error) {
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
	id := int(id64)
	return id, err
}

func (repo *UserRepository) FindById(id int) (*domain.User, error) {
	var u domain.User
	err := repo.Get(
		&u,
		"SELECT * FROM users WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (repo *UserRepository) FindAll() (domain.Users, error) {
	var users domain.Users
	err := repo.Select(
		&users,
		"SELECT * FROM users",
	)
	if err != nil {
		return nil, err
	}
	return users, nil
}
