package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
	"time"
)

func InsertNewOrderLine(orderLine *common.OrderLine) (err error) {
	if err := client.Session.Query(`INSERT INTO orderline (ol_w_id, ol_d_id, ol_o_id, ol_number, ol_supply_w_id, 
                       ol_delivery_d, ol_i_id, ol_amount, ol_quantity, ol_dist_info) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		orderLine.WarehouseID, orderLine.DistrictID, orderLine.OrderID, orderLine.ID, orderLine.SupplyWarehouseID, orderLine.DeliveryTime, orderLine.ItemID, orderLine.Amount, orderLine.Quantity, orderLine.Info).Exec(); err != nil {
		log.Printf("[warn] Create new orderline err, err=%v", err)
		return err
	}

	return nil
}

func GetOrderLineByOrder(warehouseID int32, districtID int32, orderID int32) (orderLines []common.OrderLine, err error) {
	scanner := client.Session.Query(`SELECT ol_w_id, ol_d_id, ol_o_id, ol_number, ol_i_id, ol_delivery_d, ol_amount, ol_supply_w_id, ol_quantity, ol_dist_info FROM orderline WHERE ol_w_id = ? AND ol_d_id = ? AND ol_o_id = ?`,
		warehouseID, districtID, orderID).Iter().Scanner()

	for scanner.Next() {
		orderLine := common.OrderLine{}
		err = scanner.Scan(&orderLine.WarehouseID, &orderLine.DistrictID, &orderLine.OrderID, &orderLine.ID,
			&orderLine.ItemID, &orderLine.DeliveryTime, &orderLine.Amount, &orderLine.SupplyWarehouseID,
			&orderLine.Quantity, &orderLine.Info)
		if err != nil {
			log.Printf("[warn] OrderLine info. scan error, err=%v", err)
			return []common.OrderLine{}, err
		}
		orderLines = append(orderLines, orderLine)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.OrderLine{}, err
	}

	return orderLines, nil
}

type orderLineIDItem struct {
	olNumber int32
	olItemID int32
}

func SetOrderLinesDeliveryDate(deliveryDate time.Time, warehouseID int32, districtID int32, orderID int32) (err error) {
	selectStmt := `SELECT ol_number, ol_i_id FROM orderline WHERE ol_w_id = ? AND ol_d_id = ? AND ol_o_id = ?`

	orderlines := make([]orderLineIDItem, 0)

	iter := client.Session.Query(selectStmt, warehouseID, districtID, orderID).Iter()

	for {
		var olNumber int32
		var olItemID int32
		if !iter.Scan(&olNumber, &olItemID) {
			break
		}
		orderlines = append(orderlines, orderLineIDItem{
			olNumber: olNumber,
			olItemID: olItemID,
		})
	}

	updateStmt := `UPDATE orderline SET ol_delivery_d = ? WHERE ol_w_id = ? AND ol_d_id = ? AND ol_o_id = ? AND ol_number = ? AND ol_i_id = ?`
	for _, ol := range orderlines {
		if err = client.Session.Query(updateStmt, deliveryDate, warehouseID, districtID, orderID, ol.olNumber, ol.olItemID).Exec(); err != nil {
			log.Printf("[warn] Set new carrier date err, err=%v", err)
			return err
		}
	}

	return nil
}

func GetOrderAmount(warehouseID int32, districtID int32, orderID int32) (amount float64, err error) {
	if err = client.Session.Query(`Select sum(ol_amount) from orderline where ol_w_id = ? and ol_d_id = ? and ol_o_id = ?`, warehouseID, districtID, orderID).Scan(&amount); err != nil {
		log.Printf("[warn] get order line amount err, err=%v", err)
		return 0, err
	}
	return amount, nil
}

func GetOrderIdentifiersByItemID(itemID int32) (orderIdentifiers []common.OrderIdentifierList, err error) {
	stmt := `SELECT ol_w_id, ol_d_id, ol_o_id FROM orderline WHERE ol_i_id = ?`
	scanner := client.Session.Query(stmt, itemID).Iter().Scanner()

	for scanner.Next() {
		orderIdentifier := common.OrderIdentifierList{}
		err = scanner.Scan(&orderIdentifier.WarehouseID, &orderIdentifier.DistrictID, &orderIdentifier.OrderID)
		if err != nil {
			log.Printf("[warn] OrderLine info. scan error, err=%v", err)
			return []common.OrderIdentifierList{}, err
		}
		orderIdentifiers = append(orderIdentifiers, orderIdentifier)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.OrderIdentifierList{}, err
	}

	return orderIdentifiers, nil
}
