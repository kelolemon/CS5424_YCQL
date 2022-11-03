package dao

import (
	"cs5234/client"
	"cs5234/dao"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastOrderInfo(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		res, err := dao.GetALlOrdersNotDelivery(1, 1)
		assert.NoError(t, err)
		fmt.Printf("%v", res)
	}
}
