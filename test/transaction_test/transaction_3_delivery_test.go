package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testing

func TestDelivery(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if client.Session != nil {
		defer client.Session.Close()
		resp, err := helper.CreateNewDelivery(common.CreateNewDeliveryReq{
			WarehouseID: 3,
			CarrierID:   2,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}
}
