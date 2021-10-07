package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gopla/maro/src/helper"
	"github.com/gopla/maro/src/middleware"
	"github.com/gopla/maro/src/model"
	"github.com/gopla/maro/src/service"
)

func IndexUser(c *gin.Context) {
	var user []model.User
	model.DB.Find(&user)


	res := helper.BuildResponse(true,"OK",user)

	c.JSON(http.StatusOK, res)
}

func ShowUser(c *gin.Context)  {
	var user []model.User

	if err := model.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true,"OK",user)

	c.JSON(http.StatusOK, res)
}

func StoreUser(c *gin.Context)  {
	var input model.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    c.JSON(http.StatusBadRequest, res)
    return
  }

	user := model.User{Name: input.Name, Username: input.Username, Password: helper.HashAndSalt([]byte(input.Password))}
	
	model.DB.Create(&user)

	res := helper.BuildResponse(true,"OK",user)

	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context)  {
	var user model.User

	if err := model.DB.Where("id = ?",c.Param("id")).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	
	var input model.UpdateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    c.JSON(http.StatusBadRequest, res)
    return
  }

	var updatedInput model.User
	updatedInput.Name = input.Name
	updatedInput.Username = input.Username
	updatedInput.Password = helper.HashAndSalt([]byte(input.Password))

	model.DB.Model(&user).Updates(updatedInput)
	
	res := helper.BuildResponse(true,"OK",user)

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context)  {
	var user model.User

	if err := model.DB.Where("id = ?",c.Param("id")).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
	}

	model.DB.Delete(&user)
	
	res := helper.BuildResponse(true,"OK",user)

	c.JSON(http.StatusOK, res)
}

func UpdateStatus(c *gin.Context)  {
	var userId = GetUserIdByToken(middleware.ExtractToken(c))
	var user model.User

	if err := model.DB.Where("id = ?",userId).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

 	if user.IsAvailable {
		model.DB.Model(&user).Updates(map[string]interface{}{"is_available":false})
	}else{
		model.DB.Model(&user).Updates(map[string]interface{}{"is_available":true})
	}

	res:=helper.BuildResponse(true,"UserId",user)
	c.JSON(http.StatusOK,res)
}

func GetProfile(c *gin.Context)  {
	var userId = GetUserIdByToken(middleware.ExtractToken(c))
	var user model.User

	if err := model.DB.Where("id = ?",userId).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "Profile",user)
	c.JSON(http.StatusOK,res)
}

func GetUserIdByToken(token string) string {
	aToken, err := service.JWTAuthService().ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v",claims["user_id"])
	return id
}