package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text string
	UserId uint
	Answer string
}

type CreateQuestion struct{
	Text string `json:"text" binding:"required"`
	UserId uint `json:"user_id" binding:"required"`
}

type UpdateQuestion struct{
	Text string `json:"text"`
}

type AnswerQuestion struct{
	Answer string `json:"answer"`
}