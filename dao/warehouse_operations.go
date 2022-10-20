package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func SetNewWarehouseYTD(warehouseID int32, newWarehouseYTD float64) (err error) {
	if err := client.Session.Query(`UPDATE warehouse SET w_ytd = ? WHERE w_id = ?`, newWarehouseYTD, warehouseID).Exec(); err != nil {
		log.Printf("[warn] Set new warehouse values err, err=%v", err)
		return err
	}

	return nil
}

func GetWarehouseInfo(warehouseID int32) (warehouse common.Warehouse, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM warehouse WHERE w_id = ?`, warehouseID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get warehouse information error, err=%v", err)
		return common.Warehouse{}, err
	}

	err = common.ToCqlStruct(rawMap, &warehouse)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
	}

	return warehouse, nil
}
