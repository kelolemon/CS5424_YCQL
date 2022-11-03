package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testing

func TestDelivery(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if err != nil {
		defer client.Session.Close()
		delivery, err := helper.CreateNewDelivery(common.CreateNewDeliveryReq{
			WarehouseID: 1,
			CarrierID:   5,
		})
		fmt.Printf("%v", delivery)
		assert.NoError(t, err)
	}
}
