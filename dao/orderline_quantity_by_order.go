package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func InsertOrderLineQuantityByOrderInfo(orderByCustomer *common.OrderLineQuantityByOrder) (err error) {
	err = client.Session.Query(
		`INSERT INTO OrderLineQuantityByOrder (W_ID, D_ID, O_ID, O_ENTRY_D, OL_QUANTITY_MAP, ITEMS_ID_NAME_MAP, C_FIRST, C_MIDDLE, C_LAST) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID,
		orderByCustomer.DistrictID,
		orderByCustomer.OrderID,
		orderByCustomer.OrderEntryTime,
		orderByCustomer.OrderLineQuantitiesMap,
		orderByCustomer.OrderItemsIDNameMap,
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

func GetLastLOrdersQuantity(warehouseID int32, districtID int32, numLastOrders int32) ([]common.OrderLineQuantityByOrder, error) {
	orderQuantities := make([]common.OrderLineQuantityByOrder, 0)
	stmt := `SELECT * FROM orderlinequantitybyorder WHERE w_id = ? and d_id = ? ORDER BY o_id DESC LIMIT ?`
	iter := client.Session.Query(stmt, warehouseID, districtID, numLastOrders).Iter()

	for {
		rawMap := make(map[string]interface{})
		var orderQuantity common.OrderLineQuantityByOrder
		if !iter.MapScan(rawMap) {
			break
		}
		err := common.ToCqlStruct(rawMap, &orderQuantity)
		if err != nil {
			log.Fatalf("error fetching orderQuantities: %s", err)
			return nil, err
		}
		orderQuantities = append(orderQuantities, orderQuantity)
	}

	return orderQuantities, nil
}
