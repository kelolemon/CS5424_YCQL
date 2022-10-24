package method

import (
	"cs5234/common"
	"cs5234/helper"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateNewOrder POST orders
// create a new order transaction by customer
// POST Content
// Json construct see common/type.go
func CreateNewOrder(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new order, request body = %v", string(raw))
	var createNewOrderReq common.CreateOrderReq
	err := json.Unmarshal(raw, &createNewOrderReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	res, err := helper.CreateNewOrder(createNewOrderReq)
	c.JSON(200, res)
}

// CreateNewPayment POST payments
// create a new payment
func CreateNewPayment(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new payment, request body = %v", string(raw))
	var createNewPaymentReq common.CreateNewPaymentReq
	err := json.Unmarshal(raw, &createNewPaymentReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	res, err := helper.CreateNewPayment(createNewPaymentReq)
	c.JSON(200, res)
}
