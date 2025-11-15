package controller

import (
	"net/http"
	"product-api/usecase"
	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController{
	return userController{
		UserUsecase: usecase,
	}
}

func(u *userController) GetUsers(ctx *gin.Context){
	users, err := u.UserUsecase.GetUsers()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}


