package dao

import (
	"cs5234/client"
	"log"
	"time"
)

func CreateNewOrder(orderID int32, warehouseID int32, districtID int32, customerID int32, carrierID int32, orderLineCnt int32, allLocal int32, entryDate time.Time) (err error) {
	if err := client.Session.Query(`INSERT INTO "order" (o_id, o_w_id, o_d_id, o_c_id, o_carrier_id, o_ol_cnt, 
                   o_all_local, o_entry_d) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, orderID, warehouseID, districtID,
		customerID, carrierID, orderLineCnt, allLocal, entryDate).Exec(); err != nil {
		log.Printf("[warn] Create new order err, err=%v", err)
		return err
	}

	return nil
}
