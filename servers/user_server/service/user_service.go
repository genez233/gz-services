package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func (s UserService) GetUser(id int32) {

}
