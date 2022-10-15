package method

import (
	"cs5234/common"
	"cs5234/helper"
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
	var createNewOrderReq common.CreateOrderReq
	err := json.Unmarshal(raw, &createNewOrderReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err=%v, request body =%v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	//todo create new order from db and output item info, need to add helper/CreateNewOrder and corresponding sql in dao/xx
	res, err := helper.CreateNewOrder(createNewOrderReq)
	c.JSON(200, res)
}
