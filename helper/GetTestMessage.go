package helper

import (
	"cs5234/dao"
	"github.com/gin-gonic/gin"
	"log"
)

func GetTestMessage(c *gin.Context) (msg string, err error) {
	msg, err = dao.QueryTestDBMessage()
	if err != nil {
		log.Printf("[warn] Got Msg From DB error, err=%v", err)
		return "", err
	}
	return msg, nil
}
