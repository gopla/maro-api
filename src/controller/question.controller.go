package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopla/maro/src/helper"
	"github.com/gopla/maro/src/middleware"
	"github.com/gopla/maro/src/model"
)

func IndexQuestion(c *gin.Context) {
	var question []model.Question
	model.DB.Find(&question)


	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func ShowQuestionByUsername(c *gin.Context)  {
	var question []model.Question
	var user model.User

	if err := model.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := model.DB.Where("user_id = ?", user.ID ).Find(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func ShowQuestion(c *gin.Context)  {
	var question []model.Question

	if err := model.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func StoreQuestion(c *gin.Context)  {
	var input model.CreateQuestion

	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    c.JSON(http.StatusBadRequest, res)
    return
  }

	question := model.Question{Text: input.Text, UserId: input.UserId}
	
	model.DB.Create(&question)

	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func UpdateQuestion(c *gin.Context)  {
	var question model.Question

	if err := model.DB.Where("id = ?",c.Param("id")).First(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	
	var input model.UpdateQuestion

	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    c.JSON(http.StatusBadRequest, res)
    return
  }

	var updatedInput model.Question
	updatedInput.Text = input.Text

	model.DB.Model(&question).Updates(updatedInput)
	
	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func DeleteQuestion(c *gin.Context)  {
	var question model.Question

	if err := model.DB.Where("id = ?",c.Param("id")).First(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
	}

	model.DB.Delete(&question)
	
	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func ShowQuestionPerUser(c* gin.Context)  {
	var userId = GetUserIdByToken(middleware.ExtractToken(c))
	var question []model.Question
	if err := model.DB.Where("user_id = ?", userId).Find(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}

func AnswerQuestion(c *gin.Context)  {
	var question model.Question

	if err := model.DB.Where("id = ?",c.Param("id")).First(&question).Error; err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var answer model.AnswerQuestion
	if err := c.ShouldBindJSON(&answer); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    c.JSON(http.StatusBadRequest, res)
    return
  }

	var answeredQuestion model.Question
	answeredQuestion.Answer = answer.Answer
	model.DB.Model(&question).Updates(answeredQuestion)
	
	res := helper.BuildResponse(true,"OK",question)

	c.JSON(http.StatusOK, res)
}