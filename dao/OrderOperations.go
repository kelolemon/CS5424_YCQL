package dao

import (
	"cs5234/client"
	"log"
)

func CreateNewOrder(OrderID int32, OderWarehouseID int32, OrderDistrictID int32, OrderCustomerID int32, OrderCarrierID int32, OOLCent int64, OAllLocal int32, OEntryDate int64) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return err
	}
	defer session.Close()

	if err := session.Query(`INSERT INTO Order (O_ID, O_W_ID, O_D_ID, O_C_ID, O_CARRIER_ID, O_OL_CNT, O_ALL_LOCAL, O_ENTRY_D) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		OrderID, OderWarehouseID, OrderDistrictID, OrderCustomerID, OrderCarrierID, OOLCent, OAllLocal, OEntryDate).Exec(); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return err
	}

	return nil
}
