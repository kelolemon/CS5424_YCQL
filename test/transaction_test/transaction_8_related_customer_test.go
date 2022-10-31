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
		itemQuantityMap1 := make(map[int32]int32)
		itemQuantityMap1[1] = 400
		itemQuantityMap1[2] = 20
		itemIdNameMap1 := make(map[int32]string)
		itemIdNameMap1[1] = "item1"
		itemIdNameMap1[2] = "item2"

		newOrderLineQuantity := common.OrderLineQuantityByOrder{
			WarehouseID:            1,
			DistrictID:             1,
			OrderID:                3,
			OrderEntryTime:         time.Unix(time.Now().Unix(), 0),
			OrderLineQuantitiesMap: itemQuantityMap1,
			OrderItemsIDNameMap:    itemIdNameMap1,
			CustomerFirstName:      "A",
			CustomerMiddleName:     "A",
			CustomerLastName:       "A",
		}

		err = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)
		assert.NoError(t, err)

		itemQuantityMap2 := make(map[int32]int32)
		itemQuantityMap2[10] = 5
		itemQuantityMap2[9] = 40
		itemIdNameMap2 := make(map[int32]string)
		itemIdNameMap2[10] = "item10"
		itemIdNameMap2[9] = "item9"
		newOrderLineQuantity.OrderID = 4
		newOrderLineQuantity.OrderLineQuantitiesMap = itemQuantityMap2
		newOrderLineQuantity.CustomerFirstName = "B"
		newOrderLineQuantity.CustomerMiddleName = "B"
		newOrderLineQuantity.CustomerLastName = "B"
		err = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)
		assert.NoError(t, err)

		itemQuantityMap3 := make(map[int32]int32)
		itemQuantityMap3[3] = 200
		itemQuantityMap3[4] = 50
		itemIdNameMap3 := make(map[int32]string)
		itemIdNameMap3[3] = "item3"
		itemIdNameMap3[4] = "item4"
		newOrderLineQuantity.OrderID = 5
		newOrderLineQuantity.OrderLineQuantitiesMap = itemQuantityMap3
		newOrderLineQuantity.CustomerFirstName = "C"
		newOrderLineQuantity.CustomerMiddleName = "C"
		newOrderLineQuantity.CustomerLastName = "C"
		err = dao.InsertOrderLineQuantityByOrderInfo(&newOrderLineQuantity)
		assert.NoError(t, err)
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
			CustomerID:  20,
		}

		relatedCustomerResp, _ := helper.GetRelativeCustomer(relatedCustomerReq)
		fmt.Printf("%v+\n", relatedCustomerResp)
	}
}
