package interactor

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (*domain.User, error) {
	id, err := interactor.UserRepository.Store(u)
	if err != nil {
		return nil, err
	}
	user, err := interactor.UserRepository.FindById(id)
	return user, err
}

func (interactor *UserInteractor) GetById(id int) (*domain.User, error) {
	u, err := interactor.UserRepository.FindById(id)
	return u, err
}

func (interactor *UserInteractor) GetAll() (domain.Users, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}
