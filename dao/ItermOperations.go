package dao

import (
	"cs5234/client"
	"log"
)

func GetItermPrice(itemID int32) (itemPrice float64, err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return -1, err
	}
	defer session.Close()

	if err := session.Query(`SELECT I_PRICE FROM ITEM WHERE I_ID = ?`, itemID).Scan(&itemPrice); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return 0, err
	}

	return itemPrice, nil
}
