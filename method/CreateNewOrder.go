package method

import (
	"cs5234/common"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

//CreateNewOrder POST orders
//create a new order transaction by customer
//POST Content
//Json construct see common/type.go
func CreateNewOrder(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new order, request body = %v", string(raw))
	var createNewOrderInfo common.CreateOrderInfo
	err := json.Unmarshal(raw, &createNewOrderInfo)
	if err != nil {
		log.Printf("[warn] request json converted error, err=%v, request body =%v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
