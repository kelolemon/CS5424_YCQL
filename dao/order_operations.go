package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func CreateNewOrder(order *common.Order) (err error) {
	if err := client.Session.Query(`INSERT INTO "order" (o_id, o_w_id, o_d_id, o_c_id, o_carrier_id, o_ol_cnt, 
                   o_all_local, o_entry_d) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, order.ID, order.WarehouseID, order.DistrictID,
		order.CustomerID, order.CarrierID, order.NumItemOrdered, order.OrderAllLocal, order.OrderEntryTime).Exec(); err != nil {
		log.Printf("[warn] Create new order err, err=%v", err)
		return err
	}

	return nil
}
