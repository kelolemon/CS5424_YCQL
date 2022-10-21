package dao

import (
	"cs5234/client"
	"cs5234/common"
	"fmt"
	"log"
	"time"
)

func CreateNewOrderLine(warehouseID int32, districtID int32, orderID int32, orderLineNumber int32, supplyWarehouseID int32,
	deliveryDate time.Time, itemID int32, amount float64, quantity int32, distInfo string) (err error) {
	if err := client.Session.Query(`INSERT INTO orderline (ol_w_id, ol_d_id, ol_o_id, ol_number, ol_supply_w_id, 
                       ol_delivery_d, ol_i_id, ol_amount, ol_quantity, ol_dist_info) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		warehouseID, districtID, orderID, orderLineNumber, supplyWarehouseID, deliveryDate, itemID, amount, quantity, distInfo).Exec(); err != nil {
		log.Printf("[warn] Create new orderline err, err=%v", err)
		return err
	}

	return nil
}

func GetOrderLineByOrder(orderID int32) (orderLines []common.OrderLine, err error) {
	scanner := client.Session.Query(`SELECT * FROM orderline WHERE ol_o_id = ?`, orderID).Iter().Scanner()

	var orderLine common.OrderLine
	for scanner.Next() {
		err = scanner.Scan(&orderLine.WarehouseID, &orderLine.DistrictID, &orderLine.OrderID, &orderLine.ID,
			&orderLine.SupplyWarehouseID, &orderLine.DeliveryTime, &orderLine.ItemID, &orderLine.Amount,
			&orderLine.Quantity, &orderLine.Info)
		if err != nil {
			log.Printf("[warn] OrderLine info. scan error, err=%v", err)
			return []common.OrderLine{}, err
		}
		orderLines = append(orderLines, orderLine)
		fmt.Println("OrderLine: ", orderLine)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.OrderLine{}, err
	}

	return orderLines, nil
}
