package dao

import (
	"cs5234/client"
	"github.com/yugabyte/gocql"
	"log"
)

func GetDNextOID(WareHouseID int32, DistrictID int32) (DistrictNextOrderID int32, err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return -1, err
	}
	defer session.Close()

	if err := session.Query(`SELECT D_NEXT_O_ID FROM District WHERE D_W_ID = ? and D_ID = ?`, WareHouseID, DistrictID).Consistency(gocql.One).Scan(&DistrictNextOrderID); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return 0, err
	}

	return DistrictNextOrderID, nil
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
