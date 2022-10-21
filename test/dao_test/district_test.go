package dao

import (
	"cs5234/client"
	"cs5234/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistrictSelection(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		rawMap := make(map[string]interface{})
		var district common.District
		err = client.Session.Query(`SELECT * FROM district where d_id = 1`).MapScan(rawMap)
		assert.NoError(t, err)
		err := common.ToCqlStruct(rawMap, &district)
		assert.NoError(t, err)
	}
}
