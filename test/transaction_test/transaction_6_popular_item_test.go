package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrderWithQuantity(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if client.Session != nil {
		defer client.Session.Close()
		rawMap := make(map[string]interface{})
		err := client.Session.Query(`SELECT * FROM OrderLineQuantityByOrder WHERE o_id = ?`, 3).MapScan(rawMap)
		assert.NoError(t, err)

		var orderLineQuantity common.OrderLineQuantityByOrder
		err = common.ToCqlStruct(rawMap, &orderLineQuantity)
		assert.NoError(t, err)
	}
}

func TestGetPopularItem(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if client.Session != nil {
		defer client.Session.Close()
		getPopularItemReq := common.GetPopularItemReq{
			WarehouseID:   1,
			DistrictID:    1,
			NumLastOrders: 2,
		}
		res, err := helper.GetOrderPopularItems(getPopularItemReq)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}
