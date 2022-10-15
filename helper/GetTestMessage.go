package helper

import (
	"log"

	"cs5234/dao"

	"github.com/gin-gonic/gin"
)

func GetTestMessage(c *gin.Context) (msg string, err error) {
	log.Print("[info] Query Test Message")
	msg, err = dao.QueryTestDBMessage()
	if err != nil {
		log.Printf("[warn] Got Msg From DB error, err=%v", err)
		return "", err
	}
	return msg, nil
}
