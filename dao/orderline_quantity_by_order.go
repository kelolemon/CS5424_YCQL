package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func CreateOrderLineQuantityByOrderOperationsInfo(orderByCustomer *common.OrderLineQuantityByOrder) (err error) {
	err = client.Session.Query(`INSERT INTO StockByOrderLine (W_ID, D_ID, O_ID, O_ENTRY_D, OL_QUANTITY_MAP) VALUES (?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID, orderByCustomer.DistrictID, orderByCustomer.OrderID, orderByCustomer.OrderEntryTime, orderByCustomer.OrderLineQuantitiesMap).Exec()
	if err != nil {
		log.Printf("[warn] Insert new order line quantity by order err, err=%v", err)
		return err
	}

	return nil
}
