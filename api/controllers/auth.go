package controllers

import (
	"github.com/gin-gonic/gin"
	"user-management/api/dtos/requests"
	"user-management/api/services"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	service services.AuthService
}

func New(service services.AuthService) AuthController {
	return &authController{service: service}
}

func (a authController) Register(ctx *gin.Context) {
	var user requests.UserRequest
	ctx.BindJSON(&user)
	data, err := a.service.Register(user)

	if err != nil {
		ctx.JSON(400, err)
	} else {
		ctx.JSON(201, data)
	}

}

func (a authController) Login(ctx *gin.Context) {
	var user requests.UserRequest
	ctx.BindJSON(&user)
	data, err := a.service.Login(user)

	if err != nil {
		ctx.JSON(400, err)
	} else {
		ctx.JSON(201, data)
	}
}
