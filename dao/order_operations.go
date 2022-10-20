package dao

import (
	"cs5234/client"
	"log"
	"time"
)

func CreateNewOrder(OrderID int32, OrderWarehouseID int32, OrderDistrictID int32, OrderCustomerID int32, OrderCarrierID int32, OOLCent int32, OAllLocal int32, OEntryDate time.Time) (err error) {
	if err := client.Session.Query(`INSERT INTO order (O_ID, O_W_ID, O_D_ID, O_C_ID, O_CARRIER_ID, O_OL_CNT, O_ALL_LOCAL, O_ENTRY_D) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		OrderID, OrderWarehouseID, OrderDistrictID, OrderCustomerID, OrderCarrierID, OOLCent, OAllLocal, OEntryDate).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return err
	}

	return nil
}
