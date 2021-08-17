package controllers

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/handler"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/repository"
	"github.com/ITK13201/golang-nextjs-template/backend/usecase/interactor"
)

type UserController struct {
	Interactor interactor.UserInteractor
}

func NewUserController(sqlHandler handler.SqlHandler) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: repository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
	}
}
