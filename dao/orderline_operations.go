package dao

import (
	"cs5234/client"
	"cs5234/common"
	"fmt"
	"log"
)

func CreateNewOrderLine(OLWarehouseID int32, OLDistrictID int32, OLOrderID int32, OLNumber int32, OLSupplyWarehouseID int32, OLDeliveryDate int64, OLItemID int32, OLAmount float64, OLQuantity int32, OLDistInfo string) (err error) {
	if err := client.Session.Query(`INSERT INTO orderline (OL_W_ID, OL_D_ID, OL_O_ID, OL_NUMBER, OL_SUPPLY_W_ID, OL_DELIVERY_D, OL_I_ID, OL_AMOUNT, OL_QUANTITY, OL_DIST_INFO) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		OLWarehouseID, OLDistrictID, OLOrderID, OLNumber, OLSupplyWarehouseID, OLDeliveryDate, OLItemID, OLAmount, OLQuantity, OLDistInfo).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
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
