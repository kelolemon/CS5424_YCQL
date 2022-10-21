package transaction_test

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/dao"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInsertLastOrderStatusTestData(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		// order 1 & its order lines
		currentTime := time.Unix(time.Now().Unix(), 0)
		err = dao.CreateNewOrder(1, 1, 1, 1, 1, 2, 1, currentTime)
		assert.NoError(t, err)

		err = dao.CreateNewOrderLine(1, 1, 1, 1, 1, currentTime,
			10, 5, 1, "NIL")
		assert.NoError(t, err)
		err = dao.CreateNewOrderLine(1, 1, 1, 2, 1, currentTime,
			9, 40, 2, "NIL")
		assert.NoError(t, err)

		currentOrderByCustomer := common.OrderByCustomer{
			WarehouseID:    1,
			DistrictID:     1,
			CustomerID:     1,
			OrderEntryTime: currentTime,
			FirstName:      "first",
			MiddleName:     "middle",
			LastName:       "last",
			Balance:        1000,
			CarrierID:      1,
			LastOrderID:    1,
		}
		err = dao.InsertOrderByCustomerInfo(&currentOrderByCustomer)
		assert.NoError(t, err)

		// order 2 and its order lines
		time.Sleep(2 * time.Second)
		currentTime = time.Unix(time.Now().Unix(), 0)
		err = dao.CreateNewOrder(2, 1, 1, 1, 2, 2, 1,
			time.Unix(time.Now().Unix(), 0))
		assert.NoError(t, err)

		err = dao.CreateNewOrderLine(1, 1, 2, 1, 1,
			time.Unix(time.Now().Unix(), 0), 1, 400, 1, "NIL")
		assert.NoError(t, err)
		err = dao.CreateNewOrderLine(1, 1, 2, 2, 1,
			time.Unix(time.Now().Unix(), 0), 2, 20, 2, "NIL")
		assert.NoError(t, err)

		orderByCustomer, err := dao.GetOrderByCustomerInfo(1, 1, 1)
		assert.NoError(t, err)

		err = dao.DeleteOrderByCustomerInfo(orderByCustomer.CustomerID, orderByCustomer.WarehouseID, orderByCustomer.DistrictID)
		assert.NoError(t, err)
		orderByCustomer.LastOrderID = 2
		orderByCustomer.OrderEntryTime = currentTime
		err = dao.InsertOrderByCustomerInfo(&orderByCustomer)
		assert.NoError(t, err)

		// order 3 and its order lines
		time.Sleep(2 * time.Second)
		currentTime = time.Unix(time.Now().Unix(), 0)
		err = dao.CreateNewOrder(3, 1, 1, 1, 1, 2, 1,
			time.Unix(time.Now().Unix(), 0))
		assert.NoError(t, err)

		err = dao.CreateNewOrderLine(1, 1, 3, 1, 1,
			time.Unix(time.Now().Unix(), 0), 3, 200, 1, "NIL")
		assert.NoError(t, err)
		err = dao.CreateNewOrderLine(1, 1, 3, 2, 1,
			time.Unix(time.Now().Unix(), 0), 4, 50, 2, "NIL")
		assert.NoError(t, err)

		orderByCustomer, err = dao.GetOrderByCustomerInfo(1, 1, 1)
		assert.NoError(t, err)

		err = dao.DeleteOrderByCustomerInfo(orderByCustomer.CustomerID, orderByCustomer.WarehouseID, orderByCustomer.DistrictID)
		assert.NoError(t, err)
		orderByCustomer.LastOrderID = 3
		orderByCustomer.OrderEntryTime = currentTime
		err = dao.InsertOrderByCustomerInfo(&orderByCustomer)
		assert.NoError(t, err)
	}
}

func TestGetLastOrderStatus(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		lastOrderStatusReq := common.GetLastOrderStatusReq{
			WarehouseID: 1,
			DistrictID:  1,
			CustomerID:  1,
		}

		var lastOrderStatusResp common.GetLastOrderStatusResp
		lastOrderStatusResp, err = helper.GetLastOrderStatus(lastOrderStatusReq)

		fmt.Printf("%v+\n", lastOrderStatusResp)
	}
}
