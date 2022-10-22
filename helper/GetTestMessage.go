package helper

import (
	"cs5234/common"
	"log"

	"cs5234/dao"

	"github.com/gin-gonic/gin"
)

func GetTestMessage(c *gin.Context) (msg common.DBTest, err error) {
	log.Print("[info] Query Test Message")
	msg, err = dao.QueryTestDBMessage()
	if err != nil {
		log.Printf("[warn] Got Msg From DB error, err=%v", err)
		return common.DBTest{}, err
	}
	return msg, nil
}
