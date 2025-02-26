package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello user server!")
	})

	// 启用静态文件服务
	//r.Static("/public", "./public")

	return r
}
