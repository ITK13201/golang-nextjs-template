package controllers

import (
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/handler"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/repository"
	"github.com/ITK13201/golang-nextjs-template/backend/usecase/interactor"
	"strconv"
)

type UserController struct {
	UserInteractor interactor.UserInteractor
}

func NewUserController(sqlHandler handler.SqlHandler) *UserController {
	return &UserController{
		UserInteractor: interactor.UserInteractor{
			UserRepository: repository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.UserInteractor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := controller.UserInteractor.GetById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, u)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.UserInteractor.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}
