package transaction_test

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
