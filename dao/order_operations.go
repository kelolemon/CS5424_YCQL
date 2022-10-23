package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func CreateNewOrder(order *common.Order) (err error) {
	if err = client.Session.Query(`INSERT INTO "order" (o_id, o_w_id, o_d_id, o_c_id, o_carrier_id, o_ol_cnt, 
                   o_all_local, o_entry_d) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, order.ID, order.WarehouseID, order.DistrictID,
		order.CustomerID, order.CarrierID, order.NumItemOrdered, order.OrderAllLocal, order.OrderEntryTime).Exec(); err != nil {
		log.Printf("[warn] Create new order err, err=%v", err)
		return err
	}

	return nil
}

func GetOldestNotDelivery(warehouseID int32, districtID int32) (order common.Order, err error) {
	rawMap := make(map[string]interface{})
	err = client.Session.Query(`SELECT * FROM "order" WHERE o_w_id = ? AND o_d_id = ? AND o_carrier_id = 0 order by o_id asc LIMIT 1`, warehouseID, districtID).MapScan(rawMap)
	if err != nil {
		log.Printf("[warn] Get order by customer information error, err=%v", err)
		return common.Order{}, err
	}

	err = common.ToCqlStruct(rawMap, &order)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.Order{}, err
	}

	return order, nil
}

func SetCarrierInfo(warehouseID int32, districtID int32, OrderID int32, CarrierID int32) (err error) {
	if err := client.Session.Query(`UPDATE "order" SET o_carrier_id = ? WHERE o_w_id = ? AND o_d_id = ? AND o_id = ?`, CarrierID, warehouseID, districtID, OrderID).Exec(); err != nil {
		log.Printf("[warn] Set new carrier information err, err=%v", err)
		return err
	}

	return nil
}
