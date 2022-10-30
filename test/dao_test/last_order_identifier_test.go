package dao

import (
	"cs5234/client"
	"cs5234/dao"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastOrderIdentifier(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		res, err := dao.GetOrderIdentifier(1, 1, 1)
		assert.NoError(t, err)
		fmt.Printf("%v", res)
	}
}
