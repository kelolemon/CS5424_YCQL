package transaction

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

func TestInsertCreateNewOrderTestData(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	// add district test data
	newDistrict := common.District{
		WarehouseID: 1,
		ID:          1,
		Name:        "d1",
		Street1:     "s1",
		Street2:     "s2",
		City:        "sg",
		State:       "sg",
		Zip:         "123456",
		Tax:         0,
		YTD:         0,
		NextOrderID: 3,
	}

	err = dao.InsertNewDistrictInfo(&newDistrict)
	assert.NoError(t, err)

	// add stock test data
	newStock := common.Stock{
		WarehouseID: 1,
		ItemID:      1,
		Quantity:    100,
		YTD:         50,
		OrderCnt:    1,
		RemoteCnt:   1,
		Dist1:       "NIL",
		Dist2:       "NIL",
		Dist3:       "NIL",
		Dist4:       "NIL",
		Dist5:       "NIL",
		Dist6:       "NIL",
		Dist7:       "NIL",
		Dist8:       "NIL",
		Dist9:       "NIL",
		Dist10:      "NIL",
		Data:        "NIL",
	}
	err = dao.InsertNewStockInfo(&newStock)
	assert.NoError(t, err)

	newStock.ItemID = 2
	newStock.Quantity = 10
	err = dao.InsertNewStockInfo(&newStock)
	assert.NoError(t, err)

	// add item test data
	newItem := common.Item{
		ID:      1,
		Name:    "bread",
		Price:   10,
		ImageID: 1,
		Data:    "NIL",
	}
	err = dao.InsertNewItemInfo(&newItem)
	assert.NoError(t, err)

	newItem = common.Item{
		ID:      2,
		Name:    "rib eye",
		Price:   50,
		ImageID: 2,
		Data:    "NIL",
	}
	err = dao.InsertNewItemInfo(&newItem)
	assert.NoError(t, err)

	// add customer test data
	newCustomer := common.Customer{
		WarehouseID:     1,
		DistrictID:      1,
		ID:              20,
		FirstName:       "tom",
		MiddleName:      "green",
		LastName:        "black",
		Street1:         "s1",
		Street2:         "s2",
		City:            "sg",
		State:           "sg",
		Zip:             "123456",
		Phone:           "88887777",
		CreationTime:    time.Unix(time.Now().Unix(), 0),
		CreditStatus:    "ACTIVE",
		CreditLimit:     1000,
		Discount:        0.5,
		Balance:         2000,
		YtdPayment:      500,
		NumPaymentMade:  2,
		NumDeliveryMade: 2,
		Data:            "NIL",
	}

	err = dao.InsertNewCustomerInfo(&newCustomer)
	assert.NoError(t, err)

	// add warehouse test data
	newWarehouse := common.Warehouse{
		ID:      1,
		Name:    "w1",
		Street1: "s1",
		Street2: "s2",
		City:    "sg",
		State:   "sg",
		Zip:     "123456",
		Tax:     0.2,
		YTD:     5000,
	}

	err = dao.InsertNewWarehouseInfo(&newWarehouse)
	assert.NoError(t, err)
}

func TestCreateNewOrder(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	createNewOrderReq := common.CreateOrderReq{
		WarehouseID:     1,
		DistrictID:      1,
		CustomerID:      20,
		NumberItems:     2,
		ItemNumber:      []int32{1, 2},
		SupplyWarehouse: []int32{1, 1},
		Quantity:        []int32{4, 5},
	}

	createNewOrderResp, err := helper.CreateNewOrder(createNewOrderReq)
	assert.NoError(t, err)

	fmt.Printf("%v", createNewOrderResp)
}
