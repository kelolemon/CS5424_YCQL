package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetItermInfo(itemID int32) (itemInfo common.Item, err error) {
	if err := client.Session.Query(`SELECT * FROM ITEM WHERE I_ID = ?`, itemID).Scan(&itemInfo); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return common.Item{}, err
	}

	return itemInfo, nil
}
