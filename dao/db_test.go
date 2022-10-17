package dao

import (
	"cs5234/client"
	"cs5234/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConnection(t *testing.T) {
	dao, err := NewDao(client.GetDBConfig())
	assert.NoError(t, err)

	if dao != nil {
		defer dao.Session.Close()
		var count int32
		err = dao.Session.Query(`SELECT count(*) FROM warehouse`).Scan(&count)
		assert.NoError(t, err)
	}
}

func TestWarehouseInsertion(t *testing.T) {
	dao, err := NewDao(client.GetDBConfig())
	assert.NoError(t, err)

	if dao != nil {
		defer dao.Session.Close()
		warehouse := common.Warehouse{
			ID:      1,
			Name:    "sxvnjhpd",
			Street1: "dxvcrastvybcwvmgnyk",
			Street2: "xvzxkgxtspsjdgylue",
			City:    "qflaqlocfljbepowfn",
			State:   "OM",
			Zip:     "123456789",
			Tax:     0.0384,
			Ytd:     300000.0,
		}
		err = dao.Session.Query(`INSERT INTO warehouse(w_id, w_name, w_street_1, w_street_2, w_city, w_state, w_zip, w_tax, w_ytd) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			&warehouse.ID,
			&warehouse.Name,
			&warehouse.Street1,
			&warehouse.Street2,
			&warehouse.City,
			&warehouse.State,
			&warehouse.Zip,
			&warehouse.Tax,
			&warehouse.Ytd).Exec()
		assert.NoError(t, err)
	}
}

func TestWarehouseSelection(t *testing.T) {
	dao, err := NewDao(client.GetDBConfig())
	assert.NoError(t, err)

	if dao != nil {
		defer dao.Session.Close()
		rawMap := make(map[string]interface{})
		var warehouse common.Warehouse
		err = dao.Session.Query(`SELECT * FROM warehouse where w_id = 1`).MapScan(rawMap)
		assert.NoError(t, err)
		err := common.ToCqlStruct(rawMap, &warehouse)
		assert.NoError(t, err)
	}
}
