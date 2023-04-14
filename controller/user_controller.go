package controller

import (
	"net/http"

	"github.com/febriansr/jwt-auth/model"
	"github.com/febriansr/jwt-auth/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func (c *UserController) Login(ctx *gin.Context) {
	var inputedUser model.User

	if err := ctx.ShouldBindJSON(&inputedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.Login(&inputedUser)

	if res == "user data not found" || res == "invalid credentials" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": res})
	} else if res == "failed to get user data" || res == "failed to generate token" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": res})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": res})
	}
}

func NewUserController(u usecase.UserUsecase) *UserController {
	controller := UserController{
		usecase: u,
	}
	return &controller
}
