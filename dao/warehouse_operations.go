package dao

import (
	"cs5234/client"
	"log"
)

func UpdateWarehouseYTD(warehouseID int32, Payment int32) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err,err=%v", err)
		return err
	}

	defer session.Close()

	if err := session.Query(`UPDATE Warehouse SET W_YTD = W_YTD + ? WHERE W_ID = ?`, Payment, warehouseID).Exec(); err != nil {
		log.Printf("[warn] Query err, err, err=%v", err)
		return err
	}

	return nil
}
