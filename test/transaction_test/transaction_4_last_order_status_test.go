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

// pass

func TestInsertLastOrderStatusTestData(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		// order 1 & its order lines
		currentTime := time.Unix(time.Now().Unix(), 0)
		newOrder := common.Order{
			WarehouseID:    1,
			DistrictID:     1,
			ID:             1,
			CustomerID:     1,
			CarrierID:      1,
			NumItemOrdered: 2,
			OrderAllLocal:  1,
			OrderEntryTime: currentTime,
		}

		err = dao.CreateNewOrder(&newOrder)
		assert.NoError(t, err)

		newOrderLine := common.OrderLine{
			WarehouseID:       1,
			DistrictID:        1,
			OrderID:           1,
			ID:                1,
			ItemID:            10,
			DeliveryTime:      currentTime,
			Amount:            5,
			SupplyWarehouseID: 1,
			Quantity:          1,
			Info:              "NIL",
		}

		err = dao.InsertNewOrderLine(&newOrderLine)
		assert.NoError(t, err)

		newOrderLine.ID = 9
		newOrderLine.Amount = 40
		newOrderLine.SupplyWarehouseID = 1
		newOrderLine.Quantity = 2

		err = dao.InsertNewOrderLine(&newOrderLine)
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

		newOrder.ID = 2
		newOrder.CarrierID = 2
		err = dao.CreateNewOrder(&newOrder)
		assert.NoError(t, err)

		newOrderLine.OrderID = 2
		newOrderLine.ItemID = 1
		newOrderLine.Amount = 400
		newOrderLine.Quantity = 1
		err = dao.InsertNewOrderLine(&newOrderLine)
		assert.NoError(t, err)

		newOrderLine.ID = 2
		newOrderLine.ItemID = 2
		newOrderLine.Amount = 20
		newOrderLine.Quantity = 2
		err = dao.InsertNewOrderLine(&newOrderLine)
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

		newOrder = common.Order{
			WarehouseID:    1,
			DistrictID:     1,
			ID:             3,
			CustomerID:     1,
			CarrierID:      1,
			NumItemOrdered: 2,
			OrderAllLocal:  1,
			OrderEntryTime: currentTime,
		}
		err = dao.CreateNewOrder(&newOrder)
		assert.NoError(t, err)

		newOrderLine = common.OrderLine{
			WarehouseID:       1,
			DistrictID:        1,
			OrderID:           3,
			ID:                1,
			ItemID:            3,
			DeliveryTime:      currentTime,
			Amount:            200,
			SupplyWarehouseID: 1,
			Quantity:          1,
			Info:              "NIL",
		}
		err = dao.InsertNewOrderLine(&newOrderLine)
		assert.NoError(t, err)

		newOrderLine = common.OrderLine{
			WarehouseID:       1,
			DistrictID:        1,
			OrderID:           3,
			ID:                2,
			ItemID:            4,
			DeliveryTime:      currentTime,
			Amount:            50,
			SupplyWarehouseID: 1,
			Quantity:          2,
			Info:              "NIL",
		}
		err = dao.InsertNewOrderLine(&newOrderLine)
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
