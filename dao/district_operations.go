package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetDistrictInfo(warehouseID int32, districtID int32) (district common.District, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM district WHERE d_id = ? AND d_w_id = ?`, districtID,
		warehouseID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get district information error, err=%v", err)
		return common.District{}, err
	}

	err = common.ToCqlStruct(rawMap, &district)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.District{}, err
	}

	return district, nil
}

func SetNewDistrictNextOrderID(warehouseID int32, districtID int32, districtNextOrderID int32) (err error) {
	if err := client.Session.Query(`UPDATE district SET d_next_o_id = ? WHERE d_w_id = ? AND d_id = ?`,
		districtNextOrderID, warehouseID, districtID).Exec(); err != nil {
		log.Printf("[warn] Set new district next order id err, err=%v", err)
		return err
	}

	return nil
}

func SetNewDistrictYTD(warehouseID int32, districtID int32, newDistrictYTD float64) (err error) {
	if err := client.Session.Query(`UPDATE district SET d_ytd = ? WHERE d_w_id = ? AND d_id = ?`, newDistrictYTD,
		warehouseID, districtID).Exec(); err != nil {
		log.Printf("[warn] Set new district ytd err, err=%v", err)
		return err
	}

	return nil
}

func InsertNewDistrictInfo(newDistrict *common.District) (err error) {
	err = client.Session.Query(`INSERT INTO district (d_w_id, d_id, d_zip, d_name, d_street_1, d_street_2, d_city, 
d_state, d_tax, d_ytd, d_next_o_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, newDistrict.WarehouseID, newDistrict.ID,
		newDistrict.Zip, newDistrict.Name, newDistrict.Street1, newDistrict.Street2, newDistrict.City, newDistrict.State,
		newDistrict.Tax, newDistrict.YTD, newDistrict.NextOrderID).Exec()
	if err != nil {
		log.Printf("[warn] Insert new district information err, err=%v", err)
		return err
	}

	return nil
}

func GetAllDistrictList() (districtBasicInfoLists []common.DistrictWithNameList, err error) {
	scanner := client.Session.Query(`SELECT d_w_id, d_id, d_name FROM district`).Iter().Scanner()
	districtBasicInfoList := common.DistrictWithNameList{}
	for scanner.Next() {
		err = scanner.Scan(&districtBasicInfoList.WarehouseID, &districtBasicInfoList.DistrictID, &districtBasicInfoList.DistrictName)
		if err != nil {
			log.Printf("[warn] Read district basic info list err, err=%v", err)
			return nil, err
		}
		districtBasicInfoLists = append(districtBasicInfoLists, districtBasicInfoList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []common.DistrictWithNameList{}, err
	}

	return districtBasicInfoLists, nil
}
