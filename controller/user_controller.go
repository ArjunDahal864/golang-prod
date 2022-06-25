package controller

import (
	"prod/dao"
	"prod/service"
	"prod/utils"

	"github.com/gin-gonic/gin"
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

type UserController struct {
	userService *service.UserService
}





func (uc *UserController) CreateUser(c *gin.Context) {
	var user dao.UserRequestDao
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := uc.userService.CreateUser(&user)
	if err != nil {
		utils.Response(c, 400, nil, err.Error())
		return
	}
	utils.Response(c, 200, response, "User created Successfully")
}
