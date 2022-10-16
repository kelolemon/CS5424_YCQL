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

	if session != nil {
		defer session.Close()
		var count int32
		err = session.Query(`SELECT count(*) FROM warehouse`).Consistency(gocql.One).Scan(&count)
		assert.NoError(t, err)
	}
}
