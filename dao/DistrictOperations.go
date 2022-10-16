package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetDistrictInfo(WareHouseID int32, DistrictID int32) (DistrictInfo common.District, err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return common.District{}, err
	}
	defer session.Close()

	if err := session.Query(`SELECT * FROM District WHERE D_W_ID = ? and D_ID = ?`, WareHouseID, DistrictID).Scan(&DistrictInfo); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return common.District{}, err
	}

	return DistrictInfo, nil
}

func SetNewDNextOID(WareHouseID int32, DistrictID int32, DistrictNextOrderID int32) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return err
	}
	defer session.Close()

	if err := session.Query(`UPDATE District SET D_NEXT_O_ID = ? FROM WHERE D_W_ID = ? and D_ID = ?`, DistrictNextOrderID, WareHouseID, DistrictID).Exec(); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return err
	}

	return nil
}
