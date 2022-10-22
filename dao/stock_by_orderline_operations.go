package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func CreateStockByOrderLineOperationsInfo(orderByCustomer *common.StockByOrderLine) (err error) {
	err = client.Session.Query(`INSERT INTO StockByOrderLine (W_ID, D_ID, O_ID, O_ENTRY_D, S_QUANTITY_MAP) VALUES (?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID, orderByCustomer.DistrictID, orderByCustomer.OrderID, orderByCustomer.OrderEntryTime, orderByCustomer.StockQuantitiesMap).Exec()
	if err != nil {
		log.Printf("[warn] Insert new stock by order line err, err=%v", err)
		return err
	}

	return nil
}
