package dao_test

import (
	"cs5234/client"
	"cs5234/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyOrderByCustomer(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		helper.CopyOrderByCustomer()
	}
}
