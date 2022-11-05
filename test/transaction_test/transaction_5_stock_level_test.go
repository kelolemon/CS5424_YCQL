package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockLevel(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		res, err := helper.GetStockLevelLowItemNumber(common.GetStockLevelLowItemNumberReq{
			WarehouseID:    1,
			DistrictID:     1,
			StockThreshold: 50,
			LastOrders:     3,
		})

		fmt.Printf("%v", res)
		assert.NoError(t, err)
	}
}
