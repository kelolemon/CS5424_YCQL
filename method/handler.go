package method

import (
	"cs5234/common"
	"cs5234/helper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
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
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
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
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, res)
}

// CreateNewDelivery POST delivery
// create a new delivery transaction
func CreateNewDelivery(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new delivery, request body = %v", string(raw))
	var createNewDelivery common.CreateNewDeliveryReq
	err := json.Unmarshal(raw, &createNewDelivery)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	_, err = helper.CreateNewDelivery(createNewDelivery)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "create new delivery success",
	})
}

func GetOrderStatus(c *gin.Context) {

}

// GetStockLowLevelItemNumber GET low level info
// safe action, use get method
func GetStockLowLevelItemNumber(c *gin.Context) {
	var getStockLowLevelInfo common.GetStockLevelLowItemNumberReq
	log.Printf("[info] create new delivery")
	err := c.ShouldBindQuery(&getStockLowLevelInfo)
	if err != nil {
		log.Printf("get param error, err=%v", err)
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	log.Printf("[info] Get param is %v", getStockLowLevelInfo)
	res, err := helper.GetStockLevelLowItemNumber(getStockLowLevelInfo)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, res)
}

// GetPopularItem Get most popular item for the last L orders
func GetPopularItem(c *gin.Context) {
	var getPopularItemReq common.GetPopularItemReq
	err := c.ShouldBindQuery(&getPopularItemReq)
	if err != nil {
		log.Printf("get param error, err=%v", err)
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	log.Printf("[info] getting popular items for last %d orders in warehouse %d district %d",
		getPopularItemReq.NumLastOrders, getPopularItemReq.WarehouseID, getPopularItemReq.DistrictID)
	res, err := helper.GetOrderPopularItems(getPopularItemReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, res)
}

func GetTopBalanceTransaction(c *gin.Context) {
	res, err := helper.GetTopBalanceCustomer(common.GetTopBalanceCustomerReq{})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, res)
}

func GetRelatedCustomerTransaction(c *gin.Context) {

}
