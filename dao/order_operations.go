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

func GetALlOrdersNotDelivery(warehouseID int32, districtID int32) (ordersNotDelivery []common.Order, err error) {
	scanner := client.Session.Query(`SELECT o_w_id, o_d_id, o_id, o_c_id, o_carrier_id, o_ol_cnt, o_all_local, o_entry_d FROM "order" WHERE o_w_id = ? AND o_d_id = ? AND o_carrier_id = 0`, warehouseID, districtID).Iter().Scanner()
	for scanner.Next() {
		var orderNotDelivery common.Order
		err := scanner.Scan(&orderNotDelivery.WarehouseID, &orderNotDelivery.DistrictID, &orderNotDelivery.ID, &orderNotDelivery.CustomerID,
			&orderNotDelivery.CarrierID, &orderNotDelivery.NumItemOrdered, &orderNotDelivery.OrderAllLocal, &orderNotDelivery.OrderEntryTime)

		if err != nil {
			log.Printf("[warn] Order identifiers info. scan error, err=%v", err)
			return []common.Order{}, err
		}

		ordersNotDelivery = append(ordersNotDelivery, orderNotDelivery)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.Order{}, err
	}

	return ordersNotDelivery, nil
}

func SetCarrierInfo(warehouseID int32, districtID int32, order common.Order, carrierID int32) (err error) {
	deleteStmt := `DELETE from "order" where o_w_id = ? AND o_d_id = ? AND o_id = ? AND o_carrier_id = ?`
	if err := client.Session.Query(deleteStmt, warehouseID, districtID, order.ID, 0).Exec(); err != nil {
		log.Printf("[warn] Set new carrier information err, err=%v", err)
		return err
	}
	order.CarrierID = carrierID
	if err := CreateNewOrder(&order); err != nil {
		log.Printf("[warn] Set new carrier information err, err=%v", err)
		return err
	}

	return nil
}

func GetOrderIdentifier(warehouseID int32, districtID int32, customerID int32) (orderIdentifierLists []common.OrderIdentifierList, err error) {
	scanner := client.Session.Query(`SELECT o_w_id, o_d_id, o_id FROM "order" WHERE o_w_id = ? AND o_d_id = ? AND o_c_id = ?`, warehouseID, districtID, customerID).Iter().Scanner()
	orderIdentifierList := common.OrderIdentifierList{}
	for scanner.Next() {
		err := scanner.Scan(&orderIdentifierList.WarehouseID, &orderIdentifierList.DistrictID, &orderIdentifierList.OrderID)

		if err != nil {
			log.Printf("[warn] Order identifiers info. scan error, err=%v", err)
			return []common.OrderIdentifierList{}, err
		}

		orderIdentifierLists = append(orderIdentifierLists, orderIdentifierList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.OrderIdentifierList{}, err
	}

	return orderIdentifierLists, nil
}

func GetLastOrderInfo(warehouseID int32, districtID int32, customerID int32) (orderInfo common.Order, err error) {
	scanner := client.Session.Query(`SELECT * FROM "order" WHERE o_w_id = ? AND o_d_id = ? AND o_c_id = ?`, warehouseID, districtID, customerID).Iter().Scanner()
	for scanner.Next() {
		orderInfoTemp := common.Order{}
		err := scanner.Scan(&orderInfoTemp.WarehouseID, &orderInfoTemp.DistrictID, &orderInfoTemp.ID, &orderInfoTemp.CustomerID,
			&orderInfoTemp.CarrierID, &orderInfoTemp.NumItemOrdered, &orderInfoTemp.OrderAllLocal, &orderInfoTemp.OrderEntryTime)

		if err != nil {
			log.Printf("[warn] Last order info. scan error, err=%v", err)
			return common.Order{}, err
		}

		if orderInfo.OrderEntryTime.Before(orderInfoTemp.OrderEntryTime) == true {
			orderInfo = orderInfoTemp
		}
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return common.Order{}, err
	}

	return orderInfo, nil
}

func GetAllOrderIdentifierWithDate() (orderWithDateLists []common.OrderWithDateList, err error) {
	scanner := client.Session.Query(`SELECT o_w_id, o_d_id, o_id, o_c_id, o_entry_d FROM "order"`).Iter().Scanner()
	for scanner.Next() {
		orderWithDateList := common.OrderWithDateList{}
		err := scanner.Scan(&orderWithDateList.WarehouseID, &orderWithDateList.DistrictID, &orderWithDateList.OrderID,
			&orderWithDateList.CustomerID, &orderWithDateList.OrderEntryDate)

		if err != nil {
			log.Printf("[warn] Order with date info. scan error, err=%v", err)
			return []common.OrderWithDateList{}, err
		}

		orderWithDateLists = append(orderWithDateLists, orderWithDateList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.OrderWithDateList{}, err
	}

	return orderWithDateLists, nil
}

func GetCustomerIDByOrderIdentifier(warehouseID int32, districtID int32, orderID int32) (customerID int32, err error) {
	if err := client.Session.Query(`SELECT o_c_id FROM "order" WHERE o_w_id = ? AND o_d_id = ? AND o_id = ?`, warehouseID, districtID, orderID).Scan(&customerID); err != nil {
		log.Printf("[warn] Get customerID by order identifier err, err=%v", err)
		return -1, err
	}

	return customerID, nil
}
