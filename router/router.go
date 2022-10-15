package router

import (
	"cs5234/method"

	"github.com/gin-gonic/gin"
)

func InitRouters(e *gin.Engine) {
	r := e.Group("")
	r.GET("/ping", method.Pong)
	r.GET("/test_db", method.TestDBMessage)
}
