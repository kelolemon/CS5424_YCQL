package transaction

import (
	"cs5234/client"
	"cs5234/common"
	"cs5234/dao"
	"cs5234/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// failed

func TestInsertGetTopBalanceTestData(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)
	if client.Session != nil {
		defer client.Session.Close()
		newCustomerBalance := common.CustomerBalance{
			ID:            0,
			WarehouseID:   0,
			DistrictID:    0,
			Balance:       0,
			FirstName:     "0",
			MiddleName:    "0",
			LastName:      "0",
			WarehouseName: "",
			DistrictName:  "",
		}

		for i := 10; i < 100; i += 13 {
			newCustomerBalance.ID = int32(i)
			newCustomerBalance.Balance = float64(i * 500)
			newCustomerBalance.WarehouseID = int32(i % 10)
			_ = dao.InsertCustomerBalanceInfo(&newCustomerBalance)
			assert.NoError(t, err)
		}

	}
}

func TestGetTopBalance(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		res, err := helper.GetTopBalanceCustomer(common.GetTopBalanceCustomerReq{})
		assert.NoError(t, err)

		fmt.Printf("%v", res)
	}
}
