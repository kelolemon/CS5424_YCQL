package dao

import (
	"cs5234/common"
	"log"
)

func GetDistrictsForWarehouse(WarehouseID int32) ([]common.District, error) {
	return nil, nil
}

func GetDistrictInfo(WareHouseID int32, DistrictID int32) (DistrictInfo common.District, err error) {
	if err := dao.Session.Query(`SELECT * FROM District WHERE D_W_ID = ? and D_ID = ?`, WareHouseID, DistrictID).Scan(&DistrictInfo); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return common.District{}, err
	}

	return DistrictInfo, nil
}

func SetNewDNextOID(WareHouseID int32, DistrictID int32, DistrictNextOrderID int32) (err error) {
	if err := dao.Session.Query(`UPDATE District SET D_NEXT_O_ID = ? FROM WHERE D_W_ID = ? and D_ID = ?`, DistrictNextOrderID, WareHouseID, DistrictID).Exec(); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return err
	}

	return nil
}

func SetNewDistrictYTD(warehouseID int32, districtID int32, newDistrictYTD float64) (err error) {
	if err := dao.Session.Query(`UPDATE District SET D_YTD = ? WHERE D_W_ID = ? AND D_ID = ?`, newDistrictYTD, warehouseID, districtID); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
	}

	return nil
}
