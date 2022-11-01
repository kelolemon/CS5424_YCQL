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

func TestInsertRelatedCustomerTestData(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if client.Session != nil {
		defer client.Session.Close()
		// target-order 1: {2}
		newOrder1 := common.Order{
			WarehouseID:    1,
			DistrictID:     1,
			ID:             1,
			CustomerID:     1,
			CarrierID:      0,
			NumItemOrdered: 1,
			OrderAllLocal:  1,
			OrderEntryTime: time.Unix(time.Now().Unix(), 0),
		}

		newOrderLine := common.OrderLine{
			WarehouseID:       1,
			DistrictID:        1,
			OrderID:           1,
			ID:                1,
			ItemID:            2,
			DeliveryTime:      time.Unix(time.Now().Unix(), 0),
			Amount:            5,
			SupplyWarehouseID: 0,
			Quantity:          5,
			Info:              "NIL",
		}

		_ = dao.CreateNewOrder(&newOrder1)
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLineQuantity := common.OrderLineQuantityByOrder{
			WarehouseID:            1,
			DistrictID:             1,
			OrderID:                1,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{2: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             1,
			CustomerFirstName:      "1",
			CustomerMiddleName:     "11",
			CustomerLastName:       "111",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		// target-order 2: {3, 4}
		newOrder2 := common.Order{
			WarehouseID:    1,
			DistrictID:     1,
			ID:             2,
			CustomerID:     1,
			CarrierID:      0,
			NumItemOrdered: 2,
			OrderAllLocal:  1,
			OrderEntryTime: time.Time{},
		}

		newOrderLineQuantity = common.OrderLineQuantityByOrder{
			WarehouseID:            1,
			DistrictID:             1,
			OrderID:                2,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{3: 5, 4: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             1,
			CustomerFirstName:      "1",
			CustomerMiddleName:     "11",
			CustomerLastName:       "111",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		newOrderLine.OrderID = 2
		newOrderLine.ItemID = 3

		_ = dao.CreateNewOrder(&newOrder2)
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 2
		newOrderLine.ItemID = 4
		_ = dao.InsertNewOrderLine(&newOrderLine)

		// target-order 3: {1, 5, 6}
		newOrder3 := common.Order{
			WarehouseID:    1,
			DistrictID:     1,
			ID:             3,
			CustomerID:     1,
			CarrierID:      0,
			NumItemOrdered: 3,
			OrderAllLocal:  1,
			OrderEntryTime: time.Time{},
		}

		_ = dao.CreateNewOrder(&newOrder3)
		newOrderLineQuantity = common.OrderLineQuantityByOrder{
			WarehouseID:            1,
			DistrictID:             1,
			OrderID:                3,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{1: 5, 5: 5, 6: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             1,
			CustomerFirstName:      "1",
			CustomerMiddleName:     "11",
			CustomerLastName:       "111",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		newOrderLine.ID = 1
		newOrderLine.OrderID = 3
		newOrderLine.ItemID = 1
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 2
		newOrderLine.ItemID = 5
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 3
		newOrderLine.ItemID = 6
		_ = dao.InsertNewOrderLine(&newOrderLine)

		// related-order(A) 1: {2, 3}
		newOrder4 := common.Order{
			WarehouseID:    2,
			DistrictID:     1,
			ID:             4,
			CustomerID:     2,
			CarrierID:      0,
			NumItemOrdered: 2,
			OrderAllLocal:  1,
			OrderEntryTime: time.Unix(time.Now().Unix(), 0),
		}

		_ = dao.CreateNewOrder(&newOrder4)

		newOrderLineQuantity = common.OrderLineQuantityByOrder{
			WarehouseID:            2,
			DistrictID:             1,
			OrderID:                4,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{2: 5, 3: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             2,
			CustomerFirstName:      "2",
			CustomerMiddleName:     "22",
			CustomerLastName:       "222",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		newOrderLine.WarehouseID = 2
		newOrderLine.ID = 1
		newOrderLine.OrderID = 4
		newOrderLine.ItemID = 2
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 2
		newOrderLine.ItemID = 3
		_ = dao.InsertNewOrderLine(&newOrderLine)

		// related-order(A) 2: {4}
		newOrder5 := common.Order{
			WarehouseID:    2,
			DistrictID:     1,
			ID:             5,
			CustomerID:     2,
			CarrierID:      0,
			NumItemOrdered: 1,
			OrderAllLocal:  1,
			OrderEntryTime: time.Unix(time.Now().Unix(), 0),
		}

		_ = dao.CreateNewOrder(&newOrder5)

		newOrderLineQuantity = common.OrderLineQuantityByOrder{
			WarehouseID:            2,
			DistrictID:             1,
			OrderID:                5,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{4: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             2,
			CustomerFirstName:      "2",
			CustomerMiddleName:     "22",
			CustomerLastName:       "222",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		newOrderLine.ID = 1
		newOrderLine.OrderID = 5
		newOrderLine.ItemID = 4
		_ = dao.InsertNewOrderLine(&newOrderLine)

		// related-order(B) 3: {5, 6, 7}
		newOrder6 := common.Order{
			WarehouseID:    3,
			DistrictID:     1,
			ID:             6,
			CustomerID:     3,
			CarrierID:      0,
			NumItemOrdered: 3,
			OrderAllLocal:  1,
			OrderEntryTime: time.Unix(time.Now().Unix(), 0),
		}

		_ = dao.CreateNewOrder(&newOrder6)

		newOrderLineQuantity = common.OrderLineQuantityByOrder{
			WarehouseID:            3,
			DistrictID:             1,
			OrderID:                6,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: map[int32]int32{5: 5, 6: 5, 7: 5},
			OrderItemsIDNameMap:    nil,
			CustomerID:             3,
			CustomerFirstName:      "3",
			CustomerMiddleName:     "33",
			CustomerLastName:       "333",
		}
		_ = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)

		newOrderLine.WarehouseID = 3
		newOrderLine.ID = 1
		newOrderLine.OrderID = 6
		newOrderLine.ItemID = 5
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 2
		newOrderLine.ItemID = 6
		_ = dao.InsertNewOrderLine(&newOrderLine)

		newOrderLine.ID = 3
		newOrderLine.ItemID = 7
		_ = dao.InsertNewOrderLine(&newOrderLine)
	}
}

func TestGetRelatedCustomer(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		relatedCustomerReq := common.GetRelatedCustomerReq{
			WarehouseID: 1,
			DistrictID:  1,
			CustomerID:  1,
		}

		relatedCustomerResp, _ := helper.GetRelativeCustomer(relatedCustomerReq)
		fmt.Printf("%v+\n", relatedCustomerResp)
	}
}
