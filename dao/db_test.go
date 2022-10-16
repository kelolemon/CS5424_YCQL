package dao

import (
	"cs5234/client"
	"github.com/stretchr/testify/assert"
	"github.com/yugabyte/gocql"
	"testing"
)

func TestDBConnection(t *testing.T) {
	client.InitDB()
	session, err := client.DBCluster.CreateSession()
	assert.NoError(t, err)

	defer session.Close()
	var id int32
	err = session.Query(`SELECT w_id FROM Warehouse`).Consistency(gocql.One).Scan(&id)
	assert.NoError(t, err)
}
