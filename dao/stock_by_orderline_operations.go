package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func CreateStockByOrderLineInfo(orderByCustomer *common.StockByOrderLine) (err error) {
	err = client.Session.Query(`INSERT INTO StockByOrderLine (W_ID, D_ID, O_ID, O_ENTRY_D, S_QUANTITY_MAP) VALUES (?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID, orderByCustomer.DistrictID, orderByCustomer.OrderID, orderByCustomer.OrderEntryTime, orderByCustomer.StockQuantitiesMap).Exec()
	if err != nil {
		log.Printf("[warn] Insert new stock by order line err, err=%v", err)
		return err
	}

	return nil
}

func GetLastStockByOrderLineInfo(WarehouseID int32, DistrictID int32, OrderID int32) (res common.StockByOrderLine, err error) {
	rawMap := make(map[string]interface{})
	err = client.Session.Query(`SELECT * FROM StockByOrderLine WHERE W_ID = ? AND D_ID = ? AND O_ID = ? order by O_ENTRY_D DESC LIMIT 1`,
		WarehouseID, DistrictID, OrderID).MapScan(rawMap)
	if err != nil {
		log.Printf("[warn] Get order by customer information error, err=%v", err)
		return common.StockByOrderLine{}, err
	}

	err = common.ToCqlStruct(rawMap, &res)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.StockByOrderLine{}, err
	}

	return res, nil
}
