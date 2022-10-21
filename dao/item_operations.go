package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetItermInfo(itemID int32) (item common.Item, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM item WHERE i_id = ?`, itemID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get item information error, err=%v", err)
		return common.Item{}, err
	}

	err = common.ToCqlStruct(rawMap, &item)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
	}
	return item, nil
}
