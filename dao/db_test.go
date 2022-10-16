package dao

import (
	"cs5234/client"
	"cs5234/common"
	"github.com/stretchr/testify/assert"
	"github.com/yugabyte/gocql"
	"testing"
)

func TestDBConnection(t *testing.T) {
	client.InitDB()
	session, err := client.DBCluster.CreateSession()
	assert.NoError(t, err)

	if session != nil {
		defer session.Close()
		var count int32
		err = session.Query(`SELECT count(*) FROM warehouse`).Consistency(gocql.One).Scan(&count)
		assert.NoError(t, err)
	}
}

func TestInsertion(t *testing.T) {
	client.InitDB()
	session, err := client.DBCluster.CreateSession()
	assert.NoError(t, err)

	if session != nil {
		defer session.Close()
		warehouse := common.Warehouse{
			ID:      1,
			Name:    "sxvnjhpd",
			Street1: "dxvcrastvybcwvmgnyk",
			Street2: "xvzxkgxtspsjdgylue",
			City:    "qflaqlocfljbepowfn",
			State:   "OM",
			Zip:     "123456789",
			Tax:     0.0384,

			Ytd: 300000.0,
		}
		err = session.Query(`INSERT INTO warehouse(w_id, w_name, w_street_1, w_street_2, w_city, w_state, w_zip, w_tax, w_ytd) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
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
