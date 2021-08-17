package interactor

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Insert(u)
	return err
}
