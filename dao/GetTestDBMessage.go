package dao

import (
	"cs5234/common"
	"log"

	"cs5234/client"
)

func QueryTestDBMessage() (msg common.DBTest, err error) {
	if err = client.Session.Query(`SELECT * FROM test LIMIT 1`).Scan(&msg.TestID, &msg.TestNub, &msg.TestMsg); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return common.DBTest{}, err
	}

	return msg, nil
}
