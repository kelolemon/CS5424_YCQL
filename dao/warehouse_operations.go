package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func SetNewWarehouseYTD(warehouseID int32, newWarehouseYTD float64) (err error) {
	if err := client.Session.Query(`UPDATE Warehouse SET W_YTD = ? WHERE W_ID = ?`, newWarehouseYTD, warehouseID).Exec(); err != nil {
		log.Printf("[warn] Query err, err, err=%v", err)
		return err
	}

	return nil
}

func GetWarehouseInfo(warehouseID int32) (warehouseInfo common.Warehouse, err error) {
	if err := client.Session.Query(`SELECT * FROM Warehouse WHERE W_ID = ?`, warehouseID).Scan(&warehouseInfo); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return common.Warehouse{}, err
	}

	return warehouseInfo, nil
}
