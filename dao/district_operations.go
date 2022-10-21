package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetDistrictInfo(warehouseID int32, districtID int32) (district common.District, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM district WHERE d_id = ? AND d_w_id = ?`, districtID, warehouseID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get district information error, err=%v", err)
		return common.District{}, err
	}

	err = common.ToCqlStruct(rawMap, &district)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
	}

	return district, nil
}

func SetNewDistrictNextOrderID(warehouseID int32, districtID int32, districtNextOrderID int32) (err error) {
	if err := client.Session.Query(`UPDATE district SET d_next_o_id = ? WHERE d_w_id = ? AND d_id = ?`, districtNextOrderID, warehouseID, districtID).Exec(); err != nil {
		log.Printf("[warn] Set new district next order id err, err=%v", err)
		return err
	}

	return nil
}

func SetNewDistrictYTD(warehouseID int32, districtID int32, newDistrictYTD float64) (err error) {
	if err := client.Session.Query(`UPDATE district SET d_ytd = ? WHERE d_w_id = ? AND d_id = ?`, newDistrictYTD, warehouseID, districtID).Exec(); err != nil {
		log.Printf("[warn] Set new district ytd err, err=%v", err)
		return err
	}

	return nil
}
