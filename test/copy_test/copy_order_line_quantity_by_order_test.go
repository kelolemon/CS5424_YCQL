package copy_test_test

import (
	"cs5234/client"
	"cs5234/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// deprecated

func TestCopyOrderLineQuantityByOrder(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		helper.CopyOrderLineQuantityByOrder()
	}
}
