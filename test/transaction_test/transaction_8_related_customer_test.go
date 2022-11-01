package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
