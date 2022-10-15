package dao

import (
	"cs5234/client"
	"github.com/yugabyte/gocql"
	"log"
)

func QueryTestDBMessage() (msg string, err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return "", err
	}
	defer session.Close()

	if err := session.Query(`SELECT test_msg FROM test LIMIT 1`).Consistency(gocql.One).Scan(&msg); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return "", err
	}

	return msg, nil
}
