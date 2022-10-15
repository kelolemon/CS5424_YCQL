package method

import (
	"cs5234/helper"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSONP(200, gin.H{
		"message": "pong",
	})
}

func TestDBMessage(c *gin.Context) {
	retMsg, err := helper.GetTestMessage(c)
	if err != nil {
		c.JSONP(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"message": retMsg,
	})
}
