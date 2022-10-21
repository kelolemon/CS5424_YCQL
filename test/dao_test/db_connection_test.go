package dao_test

import (
	"cs5234/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConnection(t *testing.T) {
	err := client.InitDB()
	assert.NoError(t, err)

	if client.Session != nil {
		defer client.Session.Close()
		err = client.Session.Query(`select count(*) from warehouse`).Exec()
		assert.NoError(t, err)
	}
}
