package dao

import (
	"log"

	"cs5234/client"

	"github.com/yugabyte/gocql"
)

func QueryTestDBMessage() (msg string, err error) {
	if err := client.Session.Query(`SELECT test_msg FROM test LIMIT 1`).Consistency(gocql.One).Scan(&msg); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return "", err
	}

	return msg, nil
}
