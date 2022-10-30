package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func InsertOrderLineQuantityByOrderInfo(orderByCustomer *common.OrderLineQuantityByOrder) (err error) {
	err = client.Session.Query(
		`INSERT INTO OrderLineQuantityByOrder (W_ID, D_ID, O_ID, O_ENTRY_D, OL_QUANTITY_MAP, C_FIRST, C_MIDDLE, C_LAST) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID,
		orderByCustomer.DistrictID,
		orderByCustomer.OrderID,
		orderByCustomer.OrderEntryTime,
		orderByCustomer.OrderLineQuantitiesMap,
		orderByCustomer.CustomerFirstName,
		orderByCustomer.CustomerMiddleName,
		orderByCustomer.CustomerLastName).Exec()
	if err != nil {
		log.Printf("[warn] Insert new order line quantity by order err, err=%v", err)
		return err
	}

	return nil
}

func GetOrderLineQuantity(warehouseID int32, districtID int32, orderID int32) (orderLineQuantity common.OrderLineQuantityByOrder, err error) {
	rawMap := make(map[string]interface{})
	err = client.Session.Query(`SELECT * FROM orderlinequantitybyorder WHERE w_id = ? AND d_id = ? AND o_id = ?`, warehouseID, districtID, orderID).MapScan(rawMap)
	if err != nil {
		log.Printf("[warn] Get order line quantity by order information error, err=%v", err)
		return common.OrderLineQuantityByOrder{}, err
	}

	err = common.ToCqlStruct(rawMap, &orderLineQuantity)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.OrderLineQuantityByOrder{}, err
	}

	return orderLineQuantity, nil
}
