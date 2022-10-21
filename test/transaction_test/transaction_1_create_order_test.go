package transaction

import (
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertCreateNewOrderTestData(t *testing.T) {
	
}

func TestCreateNewOrder(t *testing.T) {
	createNewOrderReq := common.CreateOrderReq{
		WarehouseID:     10,
		DistrictID:      1,
		CustomerID:      1,
		NumberItems:     2,
		ItemNumber:      []int32{1, 2},
		SupplyWarehouse: []int32{1, 1},
		Quantity:        []int32{4, 5},
	}

	createNewOrderResp, err := helper.CreateNewOrder(createNewOrderReq)
	assert.NoError(t, err)

	fmt.Printf("%v", createNewOrderResp)
}
