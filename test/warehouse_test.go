package test

import (
	"cs5234/client"
	"cs5234/common"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarehouseInsertion(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		warehouse := common.Warehouse{
			ID:      2,
			Name:    "sxvnjhpd",
			Street1: "dxvcrastvybcwvmgnyk",
			Street2: "xvzxkgxtspsjdgylue",
			City:    "qflaqlocfljbepowfn",
			State:   "OM",
			Zip:     "123456789",
			Tax:     0.0384,
			YTD:     300000.0,
		}
		err = client.Session.Query(`INSERT INTO warehouse(w_id, w_name, w_street_1, w_street_2, w_city, w_state, w_zip, w_tax, w_ytd) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			&warehouse.ID,
			&warehouse.Name,
			&warehouse.Street1,
			&warehouse.Street2,
			&warehouse.City,
			&warehouse.State,
			&warehouse.Zip,
			&warehouse.Tax,
			&warehouse.YTD).Exec()
		assert.NoError(t, err)
	}
}

func TestWarehouseSelection(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		rawMap := make(map[string]interface{})
		var warehouse common.Warehouse
		err = client.Session.Query(`SELECT * FROM warehouse where w_id = 1`).MapScan(rawMap)
		assert.NoError(t, err)
		err := common.ToCqlStruct(rawMap, &warehouse)
		assert.NoError(t, err)
		fmt.Printf("%v+\n", warehouse)
	}
}
